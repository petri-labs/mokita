package swaprouter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/petri-labs/mokita/x/swaprouter/types"
)

func (k Keeper) GetNextPoolIdAndIncrement(ctx sdk.Context) uint64 {
	return k.getNextPoolIdAndIncrement(ctx)
}

func (k Keeper) GetMokiRoutedMultihopTotalSwapFee(ctx sdk.Context, route types.MultihopRoute) (
	totalPathSwapFee sdk.Dec, sumOfSwapFees sdk.Dec, err error) {
	return k.getMokiRoutedMultihopTotalSwapFee(ctx, route)
}

// SetPoolRoutesUnsafe sets the given routes to the swaprouter keeper
// to allow routing from a pool type to a certain swap module.
// For example, balancer -> gamm.
// This utility function is only exposed for testing and should not be moved
// outside of the _test.go files.
func (k *Keeper) SetPoolRoutesUnsafe(routes map[types.PoolType]types.SwapI) {
	k.routes = routes
}
