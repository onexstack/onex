package biz

//go:generate mockgen -destination mock_biz.go -package biz onex/internal/usercenter/biz IBiz

import (
	"github.com/google/wire"
	"github.com/onexstack/onexstack/pkg/authn"

	authv1 "github.com/onexstack/onex/internal/usercenter/biz/v1/auth"
	secretv1 "github.com/onexstack/onex/internal/usercenter/biz/v1/secret"
	userv1 "github.com/onexstack/onex/internal/usercenter/biz/v1/user"
	"github.com/onexstack/onex/internal/usercenter/pkg/auth"
	"github.com/onexstack/onex/internal/usercenter/store"
)

// ProviderSet is a Wire provider set used to declare dependency injection rules.
// Includes the NewBiz constructor to create a biz instance.
// wire.Bind binds the IBiz interface to the concrete implementation *biz,
// so places that depend on IBiz will automatically inject a *biz instance.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

// IBiz defines the methods that must be implemented by the business layer.
type IBiz interface {
	// UserV1 returns the UserBiz business interface.
	UserV1() userv1.UserBiz
	// SecretV1 returns the SecretBiz business interface.
	SecretV1() secretv1.SecretBiz
	// AuthV1 returns the AuthBiz business interface.
	AuthV1() authv1.AuthBiz
}

// biz is a concrete implementation of IBiz.
type biz struct {
	store store.IStore
	authn authn.Authenticator
	auth  auth.AuthProvider
}

// Ensure that biz implements the IBiz.
var _ IBiz = (*biz)(nil)

// NewBiz creates an instance of IBiz.
func NewBiz(store store.IStore, authn authn.Authenticator, auth auth.AuthProvider) *biz {
	return &biz{store: store, authn: authn, auth: auth}
}

// UserV1 returns an instance that implements the UserBiz.
func (b *biz) UserV1() userv1.UserBiz {
	return userv1.New(b.store)
}

// SecretV1 returns an instance that implements the SecretBiz.
func (b *biz) SecretV1() secretv1.SecretBiz {
	return secretv1.New(b.store)
}

// AuthV1 returns an instance that implements the AuthBiz.
func (b *biz) AuthV1() authv1.AuthBiz {
	return authv1.New(b.store, b.authn, b.auth)
}
