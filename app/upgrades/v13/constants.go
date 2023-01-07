package v13

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/petri-labs/mokita/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Mokisis v13 upgrade.
const UpgradeName = "v13"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
