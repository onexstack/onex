package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/onexstack/onex/internal/gateway/biz IBiz

import (
	"github.com/google/wire"

	minerv1 "github.com/onexstack/onex/internal/gateway/biz/v1/miner"
	minersetv1 "github.com/onexstack/onex/internal/gateway/biz/v1/minerset"
	"github.com/onexstack/onex/internal/gateway/store"
	clientset "github.com/onexstack/onex/pkg/generated/clientset/versioned"
	"github.com/onexstack/onex/pkg/generated/informers"
)

// ProviderSet is a Wire provider set used to declare dependency injection rules.
// Includes the NewBiz constructor to create a biz instance.
// wire.Bind binds the IBiz interface to the concrete implementation *biz,
// so places that depend on IBiz will automatically inject a *biz instance.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

// IBiz defines the methods that must be implemented by the business layer.
type IBiz interface {
	// MinerSetV1 returns the MinerSetBiz business interface.
	MinerSetV1() minersetv1.MinerSetBiz
	// MinerV1 returns the MinerBiz business interface.
	MinerV1() minerv1.MinerBiz
}

// biz is a concrete implementation of IBiz.
type biz struct {
	store     store.IStore
	clientset clientset.Interface
	informer  informers.SharedInformerFactory
}

// Ensure that biz implements the IBiz.
var _ IBiz = (*biz)(nil)

// NewBiz creates an instance of IBiz.
func NewBiz(store store.IStore, clientset clientset.Interface, informer informers.SharedInformerFactory) *biz {
	return &biz{store: store, clientset: clientset, informer: informer}
}

// MinerSetV1 returns an instance that implements the MinerSetBiz.
func (b *biz) MinerSetV1() minersetv1.MinerSetBiz {
	return minersetv1.New(b.store, b.clientset, b.informer)
}

// MinerV1 returns an instance that implements the MinerBiz.
func (b *biz) MinerV1() minerv1.MinerBiz {
	return minerv1.New(b.store, b.clientset, b.informer)
}
