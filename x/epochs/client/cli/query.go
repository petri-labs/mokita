package cli

import (
	"github.com/spf13/cobra"

	"github.com/mokita-labs/mokita/mokiutils/mokicli"
	"github.com/tessornetwork/mokita/x/epochs/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := mokicli.QueryIndexCmd(types.ModuleName)
	mokicli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdEpochInfos)
	mokicli.AddQueryCmd(cmd, types.NewQueryClient, GetCmdCurrentEpoch)

	return cmd
}

func GetCmdEpochInfos() (*mokicli.QueryDescriptor, *types.QueryEpochsInfoRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "epoch-infos",
		Short: "Query running epoch infos.",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}}`,
		QueryFnName: "EpochInfos"}, &types.QueryEpochsInfoRequest{}
}

func GetCmdCurrentEpoch() (*mokicli.QueryDescriptor, *types.QueryCurrentEpochRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "current-epoch",
		Short: "Query current epoch by specified identifier.",
		Long: `{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} day`}, &types.QueryCurrentEpochRequest{}
}
