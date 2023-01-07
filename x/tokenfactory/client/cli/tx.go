package cli

import (
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/petri-labs/mokita/mokiutils/mokicli"
	"github.com/petri-labs/mokita/x/tokenfactory/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := mokicli.TxIndexCmd(types.ModuleName)
	cmd.AddCommand(
		NewCreateDenomCmd(),
		NewMintCmd(),
		NewBurnCmd(),
		// NewForceTransferCmd(),
		NewChangeAdminCmd(),
	)

	return cmd
}

func NewCreateDenomCmd() *cobra.Command {
	return mokicli.BuildTxCli[*types.MsgCreateDenom](&mokicli.TxCliDesc{
		Use:   "create-denom [subdenom] [flags]",
		Short: "create a new denom from an account. (Costs moki though!)",
	})
}

func NewMintCmd() *cobra.Command {
	return mokicli.BuildTxCli[*types.MsgMint](&mokicli.TxCliDesc{
		Use:   "mint [amount] [flags]",
		Short: "Mint a denom to an address. Must have admin authority to do so.",
	})
}

func NewBurnCmd() *cobra.Command {
	return mokicli.BuildTxCli[*types.MsgBurn](&mokicli.TxCliDesc{
		Use:   "burn [amount] [flags]",
		Short: "Burn tokens from an address. Must have admin authority to do so.",
	})
}

func NewChangeAdminCmd() *cobra.Command {
	return mokicli.BuildTxCli[*types.MsgChangeAdmin](&mokicli.TxCliDesc{
		Use:   "change-admin [denom] [new-admin-address] [flags]",
		Short: "Changes the admin address for a factory-created denom. Must have admin authority to do so.",
	})
}
