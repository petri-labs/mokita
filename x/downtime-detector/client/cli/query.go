package cli

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/petri-labs/mokita/osmoutils/mokicli"
	"github.com/petri-labs/mokita/x/downtime-detector/client/queryproto"
	"github.com/petri-labs/mokita/x/downtime-detector/types"
)

func GetQueryCmd() *cobra.Command {
	cmd := mokicli.QueryIndexCmd(types.ModuleName)
	mokicli.AddQueryCmd(cmd, queryproto.NewQueryClient, RecoveredSinceQueryCmd)

	return cmd
}

func RecoveredSinceQueryCmd() (*mokicli.QueryDescriptor, *queryproto.RecoveredSinceDowntimeOfLengthRequest) {
	return &mokicli.QueryDescriptor{
		Use:   "recovered-since downtime-duration recovery-duration",
		Short: "Queries if it has been at least <recovery-duration> since the chain was down for <downtime-duration>",
		Long: `{{.Short}}
downtime-duration is a duration, but is restricted to a smaller set. Heres a few from the set: 30s, 1m, 5m, 10m, 30m, 1h, 3 h, 6h, 12h, 24h, 36h, 48h]
{{.ExampleHeader}}
{{.CommandPrefix}} recovered-since 24h 30m`,
		CustomFieldParsers: map[string]mokicli.CustomFieldParserFn{"Downtime": parseDowntimeDuration},
	}, &queryproto.RecoveredSinceDowntimeOfLengthRequest{}
}

func parseDowntimeDuration(arg string, _ *pflag.FlagSet) (any, mokicli.FieldReadLocation, error) {
	dur, err := time.ParseDuration(arg)
	if err != nil {
		return nil, mokicli.UsedArg, err
	}
	downtime, err := types.DowntimeByDuration(dur)
	return downtime, mokicli.UsedArg, err
}
