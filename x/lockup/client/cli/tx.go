package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/osmosis-labs/osmosis/osmoutils/mokicli"
	"github.com/petri-labs/mokita/x/lockup/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := mokicli.TxIndexCmd(types.ModuleName)
	mokicli.AddTxCmd(cmd, NewLockTokensCmd)
	mokicli.AddTxCmd(cmd, NewBeginUnlockingAllCmd)
	mokicli.AddTxCmd(cmd, NewBeginUnlockByIDCmd)
	mokicli.AddTxCmd(cmd, NewForceUnlockByIdCmd)

	return cmd
}

func NewLockTokensCmd() (*mokicli.TxCliDesc, *types.MsgLockTokens) {
	return &mokicli.TxCliDesc{
		Use:   "lock-tokens [tokens]",
		Short: "lock tokens into lockup pool from user account",
		CustomFlagOverrides: map[string]string{
			"duration": FlagDuration,
		},
		Flags: mokicli.FlagDesc{RequiredFlags: []*pflag.FlagSet{FlagSetLockTokens()}},
	}, &types.MsgLockTokens{}
}

// TODO: We should change the Use string to be unlock-all
func NewBeginUnlockingAllCmd() (*mokicli.TxCliDesc, *types.MsgBeginUnlockingAll) {
	return &mokicli.TxCliDesc{
		Use:   "begin-unlock-tokens",
		Short: "begin unlock not unlocking tokens from lockup pool for sender",
	}, &types.MsgBeginUnlockingAll{}
}

// NewBeginUnlockByIDCmd unlocks individual period lock by ID.
func NewBeginUnlockByIDCmd() (*mokicli.TxCliDesc, *types.MsgBeginUnlocking) {
	return &mokicli.TxCliDesc{
		Use:   "begin-unlock-by-id [id]",
		Short: "begin unlock individual period lock by ID",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: mokicli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgBeginUnlocking{}
}

// NewForceUnlockByIdCmd force unlocks individual period lock by ID if proper permissions exist.
func NewForceUnlockByIdCmd() (*mokicli.TxCliDesc, *types.MsgForceUnlock) {
	return &mokicli.TxCliDesc{
		Use:   "force-unlock-by-id [id]",
		Short: "force unlocks individual period lock by ID",
		Long:  "force unlocks individual period lock by ID. if no amount provided, entire lock is unlocked",
		CustomFlagOverrides: map[string]string{
			"coins": FlagAmount,
		},
		Flags: mokicli.FlagDesc{OptionalFlags: []*pflag.FlagSet{FlagSetUnlockTokens()}},
	}, &types.MsgForceUnlock{}
}
