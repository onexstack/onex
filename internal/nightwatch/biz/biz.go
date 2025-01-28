package biz

//go:generate mockgen -destination mock_biz.go -package biz onex/internal/nightwatch/biz IBiz

import (
	"github.com/google/wire"

	cronjobv1 "github.com/onexstack/onex/internal/nightwatch/biz/v1/cronjob"
	jobv1 "github.com/onexstack/onex/internal/nightwatch/biz/v1/job"
	"github.com/onexstack/onex/internal/nightwatch/store"
)

// ProviderSet is a Wire provider set used to declare dependency injection rules.
// Includes the NewBiz constructor to create a biz instance.
// wire.Bind binds the IBiz interface to the concrete implementation *biz,
// so places that depend on IBiz will automatically inject a *biz instance.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

// IBiz defines the methods that must be implemented by the business layer.
type IBiz interface {
	// CronJobV1 returns the CronJobBiz business interface.
	CronJobV1() cronjobv1.CronJobBiz
	// JobV1 returns the JobBiz business interface.
	JobV1() jobv1.JobBiz
}

// biz is a concrete implementation of IBiz.
type biz struct {
	store store.IStore
}

// Ensure that biz implements the IBiz.
var _ IBiz = (*biz)(nil)

// NewBiz creates an instance of IBiz.
func NewBiz(store store.IStore) *biz {
	return &biz{store: store}
}

// CronJobV1 returns an instance that implements the CronJobBiz.
func (b *biz) CronJobV1() cronjobv1.CronJobBiz {
	return cronjobv1.New(b.store)
}

// JobV1 returns an instance that implements the JobBiz.
func (b *biz) JobV1() jobv1.JobBiz {
	return jobv1.New(b.store)
}
