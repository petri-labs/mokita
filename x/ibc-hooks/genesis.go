package ibc_hooks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"github.com/mokita-labs/mokita/mokiutils"
)

var WasmHookModuleAccountAddr sdk.AccAddress = address.Module(ModuleName, []byte("wasm-hook intermediary account"))

func IbcHooksInitGenesis(ctx sdk.Context, ak mokiutils.AccountKeeper) {
	err := mokiutils.CreateModuleAccount(ctx, ak, WasmHookModuleAccountAddr)
	if err != nil {
		panic(err)
	}
}
