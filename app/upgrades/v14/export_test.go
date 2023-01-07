package v14

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	gammkeeper "github.com/petri-labs/mokita/x/gamm/keeper"
	swaprouterkeeper "github.com/petri-labs/mokita/x/swaprouter"
)

func MigrateNextPoolId(ctx sdk.Context, gammKeeper *gammkeeper.Keeper, swaprouterKeeper *swaprouterkeeper.Keeper) {
	migrateNextPoolId(ctx, gammKeeper, swaprouterKeeper)
}
