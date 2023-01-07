package cli

import (
	"github.com/spf13/cobra"

	"github.com/petri-labs/mokita/mokiutils/mokicli"
	"github.com/petri-labs/mokita/x/concentrated-liquidity/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := mokicli.QueryIndexCmd(types.ModuleName)
	mokicli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdPool)
	mokicli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdPools)
	cmd.AddCommand(
		mokicli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
	)
	return cmd
}

func GetCmdPool() (*mokicli.QueryDescriptor, *types.QueryPoolRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "pool [poolID]",
		Short: "Query pool",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pool 1`}, &types.QueryPoolRequest{}
}

func GetCmdPools() (*mokicli.QueryDescriptor, *types.QueryPoolsRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "pools",
		Short: "Query pools",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} pools`}, &types.QueryPoolsRequest{}
}
