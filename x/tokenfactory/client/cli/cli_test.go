package cli_test

import (
	"testing"

	"github.com/petri-labs/mokita/osmoutils/mokicli"
	"github.com/petri-labs/mokita/x/tokenfactory/client/cli"
	"github.com/petri-labs/mokita/x/tokenfactory/types"
)

func TestGetCmdDenomAuthorityMetadata(t *testing.T) {
	desc, _ := cli.GetCmdDenomAuthorityMetadata()
	tcs := map[string]mokicli.QueryCliTestCase[*types.QueryDenomAuthorityMetadataRequest]{
		"basic test": {
			Cmd: "uatom",
			ExpectedQuery: &types.QueryDenomAuthorityMetadataRequest{
				Denom: "uatom",
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdDenomsFromCreator(t *testing.T) {
	desc, _ := cli.GetCmdDenomsFromCreator()
	tcs := map[string]mokicli.QueryCliTestCase[*types.QueryDenomsFromCreatorRequest]{
		"basic test": {
			Cmd: "moki1test",
			ExpectedQuery: &types.QueryDenomsFromCreatorRequest{
				Creator: "moki1test",
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}
