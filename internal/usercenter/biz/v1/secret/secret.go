package secret

//go:generate mockgen -destination mock_secret.go -package secret onex/internal/usercenter/biz/v1/secret SecretBiz

import (
	"context"
	"errors"
	"sync"

	"github.com/onexstack/onexstack/pkg/core"
	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/internal/pkg/known"
	"github.com/onexstack/onex/internal/usercenter/model"
	"github.com/onexstack/onex/internal/usercenter/pkg/conversion"
	"github.com/onexstack/onex/internal/usercenter/store"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// SecretBiz defines the interface that contains methods for handling secret requests.
type SecretBiz interface {
	// Create creates a new secret based on the provided request parameters.
	Create(ctx context.Context, rq *v1.CreateSecretRequest) (*v1.CreateSecretResponse, error)

	// Update updates an existing secret based on the provided request parameters.
	Update(ctx context.Context, rq *v1.UpdateSecretRequest) (*v1.UpdateSecretResponse, error)

	// Delete removes one or more secrets based on the provided request parameters.
	Delete(ctx context.Context, rq *v1.DeleteSecretRequest) (*v1.DeleteSecretResponse, error)

	// Get retrieves the details of a specific secret based on the provided request parameters.
	Get(ctx context.Context, rq *v1.GetSecretRequest) (*v1.GetSecretResponse, error)

	// List retrieves a list of secrets and their total count based on the provided request parameters.
	List(ctx context.Context, rq *v1.ListSecretRequest) (*v1.ListSecretResponse, error)

	// SecretExpansion defines additional methods for extended secret operations, if needed.
	SecretExpansion
}

// SecretExpansion defines additional methods for secret operations.
type SecretExpansion interface{}

// secretBiz is the implementation of the SecretBiz.
type secretBiz struct {
	store store.IStore
}

// Ensure that *secretBiz implements the SecretBiz.
var _ SecretBiz = (*secretBiz)(nil)

// New creates and returns a new instance of *secretBiz.
func New(store store.IStore) *secretBiz {
	return &secretBiz{store: store}
}

// Create implements the Create method of the SecretBiz.
func (b *secretBiz) Create(ctx context.Context, rq *v1.CreateSecretRequest) (*v1.CreateSecretResponse, error) {
	var secretM model.SecretM
	_ = core.Copy(&secretM, rq)
	// TODO: Retrieve the UserID from the custom context and assign it as needed.
	secretM.UserID = contextx.UserID(ctx)

	if err := b.store.Secret().Create(ctx, &secretM); err != nil {
		return nil, v1.ErrorSecretCreateFailed("create secret failed: %s", err.Error()) // Handle creation error.
	}

	return &v1.CreateSecretResponse{SecretID: secretM.SecretID}, nil
}

// Update implements the Update method of the SecretBiz.
func (b *secretBiz) Update(ctx context.Context, rq *v1.UpdateSecretRequest) (*v1.UpdateSecretResponse, error) {
	whr := where.T(ctx).F("name", rq.GetName())
	secretM, err := b.store.Secret().Get(ctx, whr)
	if err != nil {
		return nil, err
	}

	// Update the fields if provided in the request.
	if rq.Expires != nil {
		secretM.Expires = *rq.Expires
	}
	if rq.Status != nil {
		secretM.Status = *rq.Status
	}
	if rq.Description != nil {
		secretM.Description = *rq.Description
	}

	if err := b.store.Secret().Update(ctx, secretM); err != nil {
		return nil, err
	}

	return &v1.UpdateSecretResponse{}, nil
}

// Delete implements the Delete method of the SecretBiz.
func (b *secretBiz) Delete(ctx context.Context, rq *v1.DeleteSecretRequest) (*v1.DeleteSecretResponse, error) {
	whr := where.T(ctx).F("name", rq.GetName())
	if err := b.store.Secret().Delete(ctx, whr); err != nil {
		return nil, err
	}

	return &v1.DeleteSecretResponse{}, nil
}

// Get implements the Get method of the SecretBiz.
func (b *secretBiz) Get(ctx context.Context, rq *v1.GetSecretRequest) (*v1.GetSecretResponse, error) {
	whr := where.T(ctx).F("name", rq.GetName())
	secretM, err := b.store.Secret().Get(ctx, whr)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrorSecretNotFound(err.Error()) // Return an error if secret is not found.
		}
		return nil, err // Return any other error encountered.
	}

	return &v1.GetSecretResponse{Secret: conversion.SecretMToSecretV1(secretM)}, nil
}

// List implements the List method of the SecretBiz.
func (b *secretBiz) List(ctx context.Context, rq *v1.ListSecretRequest) (*v1.ListSecretResponse, error) {
	whr := where.T(ctx).P(int(rq.GetOffset()), int(rq.GetLimit()))
	count, secretList, err := b.store.Secret().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// Set the maximum concurrency limit using the constant MaxConcurrency
	eg.SetLimit(known.MaxErrGroupConcurrency)

	// Use goroutines to improve API performance
	for _, secret := range secretList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				converted := conversion.SecretMToSecretV1(secret)
				m.Store(secret.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.W(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	secrets := make([]*v1.Secret, 0, len(secretList))
	for _, item := range secretList {
		secret, _ := m.Load(item.ID)
		secrets = append(secrets, secret.(*v1.Secret))
	}

	return &v1.ListSecretResponse{Total: count, Secrets: secrets}, nil
}
