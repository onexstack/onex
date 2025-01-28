package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/gateway/model"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

// MinerSetMToMinerSetV1 converts a MinerSetM object from the internal model
// to a MinerSet object in the v1 API format.
func MinerSetMToMinerSetV1(minerSetModel *model.MinerSetM) *v1.MinerSet {
	var minerSet v1.MinerSet
	_ = core.CopyWithConverters(&minerSet, minerSetModel)
	return &minerSet
}

// MinerSetV1ToMinerSetM converts a MinerSet object from the v1 API format
// to a MinerSetM object in the internal model.
func MinerSetV1ToMinerSetM(minerSet *v1.MinerSet) *model.MinerSetM {
	var minerSetModel model.MinerSetM
	_ = core.CopyWithConverters(&minerSetModel, minerSet)
	return &minerSetModel
}
