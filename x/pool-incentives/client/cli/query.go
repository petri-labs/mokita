package cli

import (
	"github.com/spf13/cobra"

	"github.com/petri-labs/mokita/mokiutils/mokicli"
	"github.com/petri-labs/mokita/x/pool-incentives/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := mokicli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		GetCmdGaugeIds(),
		GetCmdDistrInfo(),
		mokicli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
		GetCmdLockableDurations(),
		GetCmdIncentivizedPools(),
		GetCmdExternalIncentiveGauges(),
	)

	return cmd
}

// GetCmdGaugeIds takes the pool id and returns the matching gauge ids and durations.
func GetCmdGaugeIds() *cobra.Command {
	return mokicli.SimpleQueryCmd[*types.QueryGaugeIdsRequest](
		"gauge-ids [pool-id]",
		"Query the matching gauge ids and durations by pool id",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} gauge-ids 1
`, types.ModuleName, types.NewQueryClient)
}

// GetCmdDistrInfo takes the pool id and returns the matching gauge ids and weights.
func GetCmdDistrInfo() *cobra.Command {
	return mokicli.SimpleQueryCmd[*types.QueryDistrInfoRequest](
		"distr-info",
		"Query distribution info",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} distr-info
`, types.ModuleName, types.NewQueryClient)
}

// GetCmdLockableDurations returns lockable durations.
func GetCmdLockableDurations() *cobra.Command {
	return mokicli.SimpleQueryCmd[*types.QueryLockableDurationsRequest](
		"lockable-durations",
		"Query lockable durations",
		`Query distribution info.

Example:
{{.CommandPrefix}} lockable-durations
`, types.ModuleName, types.NewQueryClient)
}

func GetCmdIncentivizedPools() *cobra.Command {
	return mokicli.SimpleQueryCmd[*types.QueryIncentivizedPoolsRequest](
		"incentivized-pools",
		"Query incentivized pools",
		`Query incentivized pools.

Example:
{{.CommandPrefix}} incentivized-pools
`, types.ModuleName, types.NewQueryClient)
}

func GetCmdExternalIncentiveGauges() *cobra.Command {
	return mokicli.SimpleQueryCmd[*types.QueryExternalIncentiveGaugesRequest](
		"external-incentivized-gauges",
		"Query external incentivized gauges",
		`{{.Short}}{{.ExampleHeader}}
{{.CommandPrefix}} external-incentivized-gauges
`, types.ModuleName, types.NewQueryClient)
}
