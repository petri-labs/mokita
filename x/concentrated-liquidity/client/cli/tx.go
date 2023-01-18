package cli

import (
	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/osmoutils/mokicli"
	clmodel "github.com/petri-labs/mokita/x/concentrated-liquidity/model"
	"github.com/petri-labs/mokita/x/concentrated-liquidity/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := mokicli.TxIndexCmd(types.ModuleName)
	mokicli.AddTxCmd(txCmd, NewCreatePositionCmd)
	mokicli.AddTxCmd(txCmd, NewWithdrawPositionCmd)
	mokicli.AddTxCmd(txCmd, NewCreateConcentratedPoolCmd)
	return txCmd
}

var poolIdFlagOverride = map[string]string{
	"poolid": FlagPoolId,
}

func NewCreateConcentratedPoolCmd() (*mokicli.TxCliDesc, *clmodel.MsgCreateConcentratedPool) {
	return &mokicli.TxCliDesc{
		Use:     "create-concentrated-pool [denom-0] [denom-1] [tick-spacing]",
		Short:   "create a concentrated liquidity pool with the given tick spacing",
		Example: "create-concentrated-pool uion umoki 1 --pool-id 1 --from val --chain-id mokita-1",
	}, &clmodel.MsgCreateConcentratedPool{}
}

func NewCreatePositionCmd() (*mokicli.TxCliDesc, *types.MsgCreatePosition) {
	return &mokicli.TxCliDesc{
		Use:                 "create-position [lower-tick] [upper-tick] [token-0] [token-1] [token-0-min-amount] [token-1-min-amount]",
		Short:               "create or add to existing concentrated liquidity position",
		Example:             "create-position [-69082] 69082 1000000000umoki 10000000uion 0 0 --pool-id 1 --from val --chain-id mokita-1",
		CustomFlagOverrides: poolIdFlagOverride,
		Flags:               mokicli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagSetJustPoolId()}},
	}, &types.MsgCreatePosition{}
}

func NewWithdrawPositionCmd() (*mokicli.TxCliDesc, *types.MsgWithdrawPosition) {
	return &mokicli.TxCliDesc{
		Use:                 "withdraw-position [lower-tick] [upper-tick] [liquidity-out]",
		Short:               "withdraw from an existing concentrated liquidity position",
		Example:             "withdraw-position [-69082] 69082 100317215 --pool-id 1 --from val --chain-id mokita-1",
		CustomFlagOverrides: poolIdFlagOverride,
		Flags:               mokicli.FlagDesc{RequiredFlags: []*flag.FlagSet{FlagSetJustPoolId()}},
	}, &types.MsgWithdrawPosition{}
}
