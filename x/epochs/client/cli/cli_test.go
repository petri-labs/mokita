package cli_test

import (
	"testing"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/petri-labs/mokita/x/epochs/client/cli"
	"github.com/petri-labs/mokita/x/epochs/types"
)

func TestGetCmdCurrentEpoch(t *testing.T) {
	desc, _ := cli.GetCmdCurrentEpoch()
	tcs := map[string]mokicli.QueryCliTestCase[*types.QueryCurrentEpochRequest]{
		"basic test": {
			Cmd: "day",
			ExpectedQuery: &types.QueryCurrentEpochRequest{
				Identifier: "day",
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdEpochsInfo(t *testing.T) {
	desc, _ := cli.GetCmdEpochInfos()
	tcs := map[string]mokicli.QueryCliTestCase[*types.QueryEpochsInfoRequest]{
		"basic test": {
			Cmd:           "",
			ExpectedQuery: &types.QueryEpochsInfoRequest{},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}
