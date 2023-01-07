package v14

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	ibchookstypes "github.com/petri-labs/mokita/x/ibc-hooks/types"

	"github.com/petri-labs/mokita/app/upgrades"
	cltypes "github.com/petri-labs/mokita/x/concentrated-liquidity/types"
	downtimetypes "github.com/petri-labs/mokita/x/downtime-detector/types"
	protorevtypes "github.com/petri-labs/mokita/x/protorev/types"
	swaproutertypes "github.com/petri-labs/mokita/x/swaprouter/types"
	valsetpreftypes "github.com/petri-labs/mokita/x/valset-pref/types"
)

// UpgradeName defines the on-chain upgrade name for the Mokisis v14 upgrade.
const UpgradeName = "v14"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{valsetpreftypes.StoreKey, protorevtypes.StoreKey, swaproutertypes.StoreKey, downtimetypes.StoreKey, ibchookstypes.StoreKey, cltypes.StoreKey},
		Deleted: []string{},
	},
}
