package keepers

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	transfer "github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"

	_ "github.com/petri-labs/mokita/client/docs/statik"
	concentratedliquidity "github.com/petri-labs/mokita/x/concentrated-liquidity"
	downtimemodule "github.com/petri-labs/mokita/x/downtime-detector/module"
	"github.com/petri-labs/mokita/x/epochs"
	"github.com/petri-labs/mokita/x/gamm"
	ibc_rate_limit "github.com/petri-labs/mokita/x/ibc-rate-limit"
	"github.com/petri-labs/mokita/x/incentives"
	"github.com/petri-labs/mokita/x/lockup"
	"github.com/petri-labs/mokita/x/mint"
	poolincentives "github.com/petri-labs/mokita/x/pool-incentives"
	poolincentivesclient "github.com/petri-labs/mokita/x/pool-incentives/client"
	"github.com/petri-labs/mokita/x/protorev"
	superfluid "github.com/petri-labs/mokita/x/superfluid"
	superfluidclient "github.com/petri-labs/mokita/x/superfluid/client"
	swaprouter "github.com/petri-labs/mokita/x/swaprouter/module"
	"github.com/petri-labs/mokita/x/tokenfactory"
	"github.com/petri-labs/mokita/x/twap/twapmodule"
	"github.com/petri-labs/mokita/x/txfees"
	valsetprefmodule "github.com/petri-labs/mokita/x/valset-pref/valpref-module"
	ibc_hooks "github.com/petri-labs/mokita/x/ibc-hooks"
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = []module.AppModuleBasic{
	auth.AppModuleBasic{},
	genutil.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	downtimemodule.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(
		append(
			wasmclient.ProposalHandlers,
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			poolincentivesclient.UpdatePoolIncentivesHandler,
			poolincentivesclient.ReplacePoolIncentivesHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
			superfluidclient.SetSuperfluidAssetsProposalHandler,
			superfluidclient.RemoveSuperfluidAssetsProposalHandler,
			superfluidclient.UpdateUnpoolWhitelistProposalHandler,
		)...,
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	gamm.AppModuleBasic{},
	swaprouter.AppModuleBasic{},
	twapmodule.AppModuleBasic{},
	concentratedliquidity.AppModuleBasic{},
	protorev.AppModuleBasic{},
	txfees.AppModuleBasic{},
	incentives.AppModuleBasic{},
	lockup.AppModuleBasic{},
	poolincentives.AppModuleBasic{},
	epochs.AppModuleBasic{},
	superfluid.AppModuleBasic{},
	tokenfactory.AppModuleBasic{},
	valsetprefmodule.AppModuleBasic{},
	wasm.AppModuleBasic{},
	ica.AppModuleBasic{},
	ibc_hooks.AppModuleBasic{},
	ibc_rate_limit.AppModuleBasic{},
}
