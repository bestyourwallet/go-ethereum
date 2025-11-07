package arbitrum

import (
	"context"

	"github.com/bestyourwallet/go-ethereum/common/hexutil"
	"github.com/bestyourwallet/go-ethereum/core"
	"github.com/bestyourwallet/go-ethereum/internal/ethapi"
	"github.com/bestyourwallet/go-ethereum/internal/ethapi/override"
	"github.com/bestyourwallet/go-ethereum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *override.StateOverride, blockOverrides *override.BlockOverrides, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, blockOverrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}
