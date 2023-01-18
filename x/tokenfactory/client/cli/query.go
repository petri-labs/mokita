package cli

import (

	// "strings"

	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mokita-labs/mokita/mokiutils/mokicli"
	"github.com/tessornetwork/mokita/x/tokenfactory/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := mokicli.QueryIndexCmd(types.ModuleName)

	mokicli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdDenomAuthorityMetadata)
	mokicli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdDenomAuthorityMetadata)

	cmd.AddCommand(
		mokicli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
	)

	return cmd
}

func GetCmdDenomAuthorityMetadata() (*mokicli.QueryDescriptor, *types.QueryDenomAuthorityMetadataRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "denom-authority-metadata [denom] [flags]",
		Short: "Get the authority metadata for a specific denom",
		Long: `{{.Short}}{{.ExampleHeader}}
		{{.CommandPrefix}} uatom`,
	}, &types.QueryDenomAuthorityMetadataRequest{}
}

func GetCmdDenomsFromCreator() (*mokicli.QueryDescriptor, *types.QueryDenomsFromCreatorRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "denoms-from-creator [creator address] [flags]",
		Short: "Returns a list of all tokens created by a specific creator address",
		Long: `{{.Short}}{{.ExampleHeader}}
		{{.CommandPrefix}} <address>`,
	}, &types.QueryDenomsFromCreatorRequest{}
}
