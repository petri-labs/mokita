package cli

import (
	"github.com/spf13/cobra"

	"github.com/petri-labs/mokita/osmoutils/mokicli"
	"github.com/petri-labs/mokita/x/ibc-rate-limit/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := mokicli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		mokicli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
	)

	return cmd
}
