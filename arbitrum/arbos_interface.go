package arbitrum

import (
	"context"

	"github.com/bestyourwallet/go-ethereum/arbitrum_types"
	"github.com/bestyourwallet/go-ethereum/core"
	"github.com/bestyourwallet/go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
