package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	epochstypes "github.com/petri-labs/mokita/x/epochs/types"
	"github.com/petri-labs/mokita/x/protorev/types"
)

type EpochHooks struct {
	k Keeper
}

// Struct used to track the pool with the highest liquidity
type LiquidityPoolStruct struct {
	Liquidity sdk.Int
	PoolId    uint64
}

var _ epochstypes.EpochHooks = EpochHooks{}

func (k Keeper) EpochHooks() epochstypes.EpochHooks {
	return EpochHooks{k}
}

// BeforeEpochStart is the epoch start hook.
func (h EpochHooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return nil
}

// AfterEpochEnd is the epoch end hook. The module will update all of the pools in the store that are
// used for trading.
func (h EpochHooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	switch epochIdentifier {
	case "week":
		// Update the pools in the store
		return h.k.UpdatePools(ctx)
	}
	return nil
}

// Update pools requests the highest liquidity pools from the gamm module and updates the pools in the store
func (k Keeper) UpdatePools(ctx sdk.Context) error {
	// Reset the pools in the store
	k.DeleteAllAtomPools(ctx)
	k.DeleteAllMokiPools(ctx)

	// Get the highest liquidity pools
	mokiPools, atomPools, err := k.GetHighestLiquidityPools(ctx)
	if err != nil {
		return err
	}

	// Update the pools in the store
	for token, poolInfo := range mokiPools {
		k.SetMokiPool(ctx, token, poolInfo.PoolId)
	}
	for token, poolInfo := range atomPools {
		k.SetAtomPool(ctx, token, poolInfo.PoolId)
	}

	return nil
}

// GetHighestLiquidityPools returns the highest liquidity pools for pools that have Moki or Atom
// and Moki/Atom
func (k Keeper) GetHighestLiquidityPools(ctx sdk.Context) (map[string]LiquidityPoolStruct, map[string]LiquidityPoolStruct, error) {
	// Get all pools
	pools, err := k.gammKeeper.GetPoolsAndPoke(ctx)
	if err != nil {
		return nil, nil, err
	}

	mokiPools := make(map[string]LiquidityPoolStruct)
	atomPools := make(map[string]LiquidityPoolStruct)

	// Iterate through all pools and find valid matches
	for _, pool := range pools {
		coins := pool.GetTotalPoolLiquidity(ctx)

		// Pool must be active and the number of coins must be 2
		if pool.IsActive(ctx) && len(coins) == 2 {
			tokenA := coins[0]
			tokenB := coins[1]

			newPool := LiquidityPoolStruct{
				PoolId:    pool.GetId(),
				Liquidity: tokenA.Amount.Mul(tokenB.Amount),
			}

			// Check if there is a match with moki
			if otherDenom, match := types.CheckMokiAtomDenomMatch(tokenA.Denom, tokenB.Denom, types.MokisisDenomination); match {
				k.updateHighestLiquidityPool(otherDenom, mokiPools, newPool)
			}

			// Check if there is a match with atom
			if otherDenom, match := types.CheckMokiAtomDenomMatch(tokenA.Denom, tokenB.Denom, types.AtomDenomination); match {
				k.updateHighestLiquidityPool(otherDenom, atomPools, newPool)
			}
		}
	}

	return mokiPools, atomPools, nil
}

// updateHighestLiquidityPool updates the pool with the highest liquidity for either moki or atom
func (k Keeper) updateHighestLiquidityPool(denom string, pool map[string]LiquidityPoolStruct, newPool LiquidityPoolStruct) {
	if currPool, ok := pool[denom]; !ok {
		pool[denom] = newPool
	} else {
		if newPool.Liquidity.GT(currPool.Liquidity) {
			pool[denom] = newPool
		}
	}
}
