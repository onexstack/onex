package miner

//go:generate mockgen -destination mock_miner.go -package miner onex/internal/gateway/biz/v1/miner MinerBiz

import (
	"context"
	"sync"

	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/store/where"
	"golang.org/x/sync/errgroup"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	"github.com/onexstack/onex/internal/gateway/pkg/conversion"
	"github.com/onexstack/onex/internal/gateway/store"
	"github.com/onexstack/onex/internal/pkg/contextx"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	clientset "github.com/onexstack/onex/pkg/generated/clientset/versioned"
	"github.com/onexstack/onex/pkg/generated/informers"
	listers "github.com/onexstack/onex/pkg/generated/listers/apps/v1beta1"
)

const (
	// MaxErrGroupConcurrency defines the maximum concurrency level
	// for error group operations.
	MaxErrGroupConcurrency = 100
)

// MinerBiz defines the interface that contains methods for handling miner requests.
type MinerBiz interface {
	// Create creates a new miner based on the provided request parameters.
	Create(ctx context.Context, miner *v1beta1.Miner) (*v1beta1.Miner, error)

	// Update updates an existing miner based on the provided request parameters.
	Update(ctx context.Context, miner *v1beta1.Miner) (*v1beta1.Miner, error)

	// Delete removes one or more miners based on the provided request parameters.
	Delete(ctx context.Context, rq *v1.DeleteMinerRequest) (*v1.DeleteMinerResponse, error)

	// Get retrieves the details of a specific miner based on the provided request parameters.
	Get(ctx context.Context, rq *v1.GetMinerRequest) (*v1beta1.Miner, error)

	// List retrieves a list of miners and their total count based on the provided request parameters.
	List(ctx context.Context, rq *v1.ListMinerRequest) (*v1.ListMinerResponse, error)

	// MinerExpansion defines additional methods for extended miner operations, if needed.
	MinerExpansion
}

// MinerExpansion defines additional methods for miner operations.
type MinerExpansion interface{}

// minerBiz is the implementation of the MinerBiz.
type minerBiz struct {
	store     store.IStore        // Data store interface for accessing miner data.
	clientset clientset.Interface // Kubernetes client interface for interacting with the API.
	lister    listers.MinerLister // Lister interface for retrieving miners from the cache.
}

// Ensure that *minerBiz implements the MinerBiz.
var _ MinerBiz = (*minerBiz)(nil)

// New creates and returns a new instance of *minerBiz.
func New(store store.IStore, clientset clientset.Interface, informer informers.SharedInformerFactory) *minerBiz {
	return &minerBiz{
		store:     store,
		clientset: clientset,
		lister:    informer.Apps().V1beta1().Miners().Lister(),
	}
}

// Create implements the Create method of the MinerBiz.
func (b *minerBiz) Create(ctx context.Context, miner *v1beta1.Miner) (*v1beta1.Miner, error) {
	return b.clientset.AppsV1beta1().Miners(contextx.Namespace(ctx)).Create(ctx, miner, metav1.CreateOptions{})
}

// Update implements the Update method of the MinerBiz.
func (b *minerBiz) Update(ctx context.Context, miner *v1beta1.Miner) (*v1beta1.Miner, error) {
	return b.clientset.AppsV1beta1().Miners(contextx.Namespace(ctx)).Update(ctx, miner, metav1.UpdateOptions{})
}

// Delete implements the Delete method of the MinerBiz.
func (b *minerBiz) Delete(ctx context.Context, rq *v1.DeleteMinerRequest) (*v1.DeleteMinerResponse, error) {
	namespace := contextx.UserID(ctx)
	if err := b.clientset.AppsV1beta1().Miners(namespace).Delete(ctx, rq.Name, metav1.DeleteOptions{}); err != nil {
		log.W(ctx).Errorw(err, "Failed to delete miner", "miner", klog.KRef(namespace, rq.Name))
		return nil, err
	}

	return &v1.DeleteMinerResponse{}, nil
}

// Get implements the Get method of the MinerBiz.
func (b *minerBiz) Get(ctx context.Context, rq *v1.GetMinerRequest) (*v1beta1.Miner, error) {
	return b.lister.Miners(contextx.Namespace(ctx)).Get(rq.Name)
}

// List implements the List method of the MinerBiz.
func (b *minerBiz) List(ctx context.Context, rq *v1.ListMinerRequest) (*v1.ListMinerResponse, error) {
	whr := where.F("namespace", contextx.Namespace(ctx)).P(int(rq.Offset), int(rq.Limit))
	count, minerList, err := b.store.Miner().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// Set the maximum concurrency limit using the constant MaxConcurrency
	eg.SetLimit(MaxErrGroupConcurrency)

	// Use goroutines to improve API performance
	for _, miner := range minerList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				converted := conversion.MinerMToMinerV1(miner)
				m.Store(miner.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.W(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	miners := make([]*v1.Miner, 0, len(minerList))
	for _, item := range minerList {
		miner, _ := m.Load(item.ID)
		miners = append(miners, miner.(*v1.Miner))
	}

	return &v1.ListMinerResponse{Total: count, Miners: miners}, nil
}
