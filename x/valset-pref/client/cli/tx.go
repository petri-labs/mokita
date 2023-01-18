package valsetprefcli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/petri-labs/mokita/osmoutils"
	"github.com/petri-labs/mokita/osmoutils/mokicli"
	"github.com/petri-labs/mokita/x/valset-pref/types"
)

func GetTxCmd() *cobra.Command {
	txCmd := mokicli.TxIndexCmd(types.ModuleName)
	mokicli.AddTxCmd(txCmd, NewSetValSetCmd)
	mokicli.AddTxCmd(txCmd, NewDelValSetCmd)
	mokicli.AddTxCmd(txCmd, NewUnDelValSetCmd)
	return txCmd
}

func NewSetValSetCmd() (*mokicli.TxCliDesc, *types.MsgSetValidatorSetPreference) {
	return &mokicli.TxCliDesc{
		Use:              "set-valset [delegator_addr] [validators] [weights]",
		Short:            "Creates a new validator set for the delegator with valOperAddress and weight",
		Example:          "mokitad tx valset-pref set-valset moki1... mokivaloper1abc...,mokivaloper1def...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgSetValidatorSetPreference,
	}, &types.MsgSetValidatorSetPreference{}
}

func NewDelValSetCmd() (*mokicli.TxCliDesc, *types.MsgDelegateToValidatorSet) {
	return &mokicli.TxCliDesc{
		Use:     "delegate-valset [delegator_addr] [amount]",
		Short:   "Delegate tokens to existing valset using delegatorAddress and tokenAmount.",
		Example: "mokitad tx valset-pref delegate-valset  moki1... 100stake",
		NumArgs: 2,
	}, &types.MsgDelegateToValidatorSet{}
}

func NewUnDelValSetCmd() (*mokicli.TxCliDesc, *types.MsgUndelegateFromValidatorSet) {
	return &mokicli.TxCliDesc{
		Use:     "undelegate-valset [delegator_addr] [amount]",
		Short:   "UnDelegate tokens from existing valset using delegatorAddress and tokenAmount.",
		Example: "mokitad tx valset-pref undelegate-valset  moki1... 100stake",
		NumArgs: 2,
	}, &types.MsgUndelegateFromValidatorSet{}
}

func NewMsgSetValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	var valAddrs []string
	valAddrs = append(valAddrs, strings.Split(args[1], ",")...)

	weights, err := osmoutils.ParseSdkDecFromString(args[2], ",")
	if err != nil {
		return nil, err
	}

	if len(valAddrs) != len(weights) {
		return nil, fmt.Errorf("the length of validator addresses and weights not matched")
	}

	if len(valAddrs) == 0 {
		return nil, fmt.Errorf("records is empty")
	}

	var valset []types.ValidatorPreference
	for i, val := range valAddrs {
		valset = append(valset, types.ValidatorPreference{
			ValOperAddress: val,
			Weight:         weights[i],
		})
	}

	return types.NewMsgSetValidatorSetPreference(
		delAddr,
		valset,
	), nil
}
