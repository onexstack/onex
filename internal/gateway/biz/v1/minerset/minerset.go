package minerset

//go:generate mockgen -destination mock_minerset.go -package minerset github.com/onexstack/onex/internal/gateway/biz/v1/minerset MinerSetBiz

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
	"github.com/onexstack/onex/internal/pkg/known"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
	"github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	clientset "github.com/onexstack/onex/pkg/generated/clientset/versioned"
	"github.com/onexstack/onex/pkg/generated/informers"
	listers "github.com/onexstack/onex/pkg/generated/listers/apps/v1beta1"
)

// MinerSetBiz defines the interface that contains methods for handling minerset requests.
type MinerSetBiz interface {
	// Create creates a new minerset based on the provided request parameters.
	Create(ctx context.Context, minerSet *v1beta1.MinerSet) (*v1beta1.MinerSet, error)

	// Update updates an existing minerset based on the provided request parameters.
	Update(ctx context.Context, minerSet *v1beta1.MinerSet) (*v1beta1.MinerSet, error)

	// Delete removes one or more minersets based on the provided request parameters.
	Delete(ctx context.Context, rq *v1.DeleteMinerSetRequest) (*v1.DeleteMinerSetResponse, error)

	// Get retrieves the details of a specific minerset based on the provided request parameters.
	Get(ctx context.Context, rq *v1.GetMinerSetRequest) (*v1beta1.MinerSet, error)

	// List retrieves a list of minersets and their total count based on the provided request parameters.
	List(ctx context.Context, rq *v1.ListMinerSetRequest) (*v1.ListMinerSetResponse, error)

	// MinerSetExpansion defines additional methods for extended minerset operations, if needed.
	MinerSetExpansion
}

// MinerSetExpansion defines additional methods for minerset operations.
type MinerSetExpansion interface {
	// Scale adjusts the number of replicas for a miner set.
	Scale(ctx context.Context, rq *v1.ScaleMinerSetRequest) (*v1.ScaleMinerSetResponse, error)
}

// minerSetBiz is the implementation of the MinerSetBiz.
type minerSetBiz struct {
	store     store.IStore           // Data store interface for accessing miner data.
	clientset clientset.Interface    // Kubernetes client interface for interacting with the API.
	lister    listers.MinerSetLister // Lister interface for retrieving miner sets from the cache.
}

// Ensure that *minerSetBiz implements the MinerSetBiz.
var _ MinerSetBiz = (*minerSetBiz)(nil)

// New creates and returns a new instance of *minerSetBiz.
func New(store store.IStore, clientset clientset.Interface, informer informers.SharedInformerFactory) *minerSetBiz {
	return &minerSetBiz{
		store:     store,
		clientset: clientset,
		lister:    informer.Apps().V1beta1().MinerSets().Lister(),
	}
}

// Create implements the Create method of the MinerSetBiz.
func (b *minerSetBiz) Create(ctx context.Context, minerSet *v1beta1.MinerSet) (*v1beta1.MinerSet, error) {
	return b.clientset.AppsV1beta1().MinerSets(contextx.Namespace(ctx)).Create(ctx, minerSet, metav1.CreateOptions{})
}

// Update implements the Update method of the MinerSetBiz.
func (b *minerSetBiz) Update(ctx context.Context, minerSet *v1beta1.MinerSet) (*v1beta1.MinerSet, error) {
	return b.clientset.AppsV1beta1().MinerSets(contextx.Namespace(ctx)).Update(ctx, minerSet, metav1.UpdateOptions{})
}

// Delete implements the Delete method of the MinerSetBiz.
func (b *minerSetBiz) Delete(ctx context.Context, rq *v1.DeleteMinerSetRequest) (*v1.DeleteMinerSetResponse, error) {
	namespace := contextx.Namespace(ctx)
	if err := b.clientset.AppsV1beta1().MinerSets(namespace).Delete(ctx, rq.Name, metav1.DeleteOptions{}); err != nil {
		log.W(ctx).Errorw(err, "Failed to delete miner set", "minerset", klog.KRef(namespace, rq.Name))
		return nil, err
	}

	return &v1.DeleteMinerSetResponse{}, nil
}

// Get implements the Get method of the MinerSetBiz.
func (b *minerSetBiz) Get(ctx context.Context, rq *v1.GetMinerSetRequest) (*v1beta1.MinerSet, error) {
	return b.lister.MinerSets(contextx.Namespace(ctx)).Get(rq.Name)
}

// List implements the List method of the MinerSetBiz.
func (b *minerSetBiz) List(ctx context.Context, rq *v1.ListMinerSetRequest) (*v1.ListMinerSetResponse, error) {
	whr := where.F("namespace", contextx.Namespace(ctx)).P(int(rq.Offset), int(rq.Limit))
	count, minerSetList, err := b.store.MinerSet().List(ctx, whr)
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)

	// Set the maximum concurrency limit using the constant MaxConcurrency
	eg.SetLimit(known.MaxErrGroupConcurrency)

	// Use goroutines to improve API performance
	for _, minerSet := range minerSetList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				converted := conversion.MinerSetMToMinerSetV1(minerSet)
				m.Store(minerSet.ID, converted)

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.W(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	minerSets := make([]*v1.MinerSet, 0, len(minerSetList))
	for _, item := range minerSetList {
		minerSet, _ := m.Load(item.ID)
		minerSets = append(minerSets, minerSet.(*v1.MinerSet))
	}

	return &v1.ListMinerSetResponse{Total: count, MinerSets: minerSets}, nil
}

// Scale adjusts the number of replicas for a miner set.
func (b *minerSetBiz) Scale(ctx context.Context, rq *v1.ScaleMinerSetRequest) (*v1.ScaleMinerSetResponse, error) {
	namespace := contextx.Namespace(ctx)
	// Retrieve the current scale configuration for the miner set.
	scale, err := b.clientset.AppsV1beta1().MinerSets(namespace).GetScale(ctx, rq.Name, metav1.GetOptions{})
	if err != nil {
		log.W(ctx).Errorw(err, "Failed to get scale", "minerset", klog.KRef(namespace, rq.Name))
		return nil, err
	}

	// Update the replicas count in the scale specification.
	scale.Spec.Replicas = rq.Replicas
	if _, err := b.clientset.AppsV1beta1().MinerSets(namespace).UpdateScale(ctx, rq.Name, scale, metav1.UpdateOptions{}); err != nil {
		log.W(ctx).Errorw(err, "Failed to scale miner set", "minerset", klog.KRef(namespace, rq.Name))
		return &v1.ScaleMinerSetResponse{}, err
	}

	return &v1.ScaleMinerSetResponse{}, nil
}
