package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	swaproutertypes "github.com/petri-labs/mokita/x/swaprouter/types"
)

// BankKeeper defines the banking contract that must be fulfilled when
// creating a x/concentrated-liquidity keeper.
type BankKeeper interface {
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// SwaprouterKeeper defines the interface needed to be fulfilled for
// the swaprouter keeper.
type SwaprouterKeeper interface {
	CreatePool(ctx sdk.Context, msg swaproutertypes.CreatePoolMsg) (uint64, error)
	GetNextPoolId(ctx sdk.Context) uint64
}
