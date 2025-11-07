package eth

import (
	"context"

	"github.com/bestyourwallet/go-ethereum/core"
	"github.com/bestyourwallet/go-ethereum/core/state"
	"github.com/bestyourwallet/go-ethereum/core/types"
	"github.com/bestyourwallet/go-ethereum/core/vm"
	"github.com/bestyourwallet/go-ethereum/eth/tracers"
	"github.com/bestyourwallet/go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*types.Transaction, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
