package conversion

import (
	"github.com/onexstack/onexstack/pkg/core"

	"github.com/onexstack/onex/internal/gateway/model"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

// MinerMToMinerV1 converts a MinerM object from the internal model
// to a Miner object in the v1 API format.
func MinerMToMinerV1(minerModel *model.MinerM) *v1.Miner {
	var miner v1.Miner
	_ = core.CopyWithConverters(&miner, minerModel)
	return &miner
}

// MinerV1ToMinerM converts a Miner object from the v1 API format
// to a MinerM object in the internal model.
func MinerV1ToMinerM(miner *v1.Miner) *model.MinerM {
	var minerModel model.MinerM
	_ = core.CopyWithConverters(&minerModel, miner)
	return &minerModel
}
