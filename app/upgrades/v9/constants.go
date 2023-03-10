package v9

import (
	"github.com/petri-labs/mokita/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"

	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"

	tokenfactorytypes "github.com/petri-labs/mokita/x/tokenfactory/types"
)

// UpgradeName defines the on-chain upgrade name for the Mokita v9 upgrade.
const UpgradeName = "v9"

// The historic name of the claims module, which is removed in this release.
// Cross-check against https://github.com/petri-labs/mokita/blob/v7.2.0/x/claim/types/keys.go#L5
const ClaimsModuleName = "claim"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{tokenfactorytypes.ModuleName, icahosttypes.StoreKey},
		Deleted: []string{ClaimsModuleName},
	},
}
