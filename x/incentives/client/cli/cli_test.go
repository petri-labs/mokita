package cli

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/petri-labs/mokita/mokiutils"
	"github.com/petri-labs/mokita/mokiutils/mokicli"
	"github.com/petri-labs/mokita/x/incentives/types"
)

var testAddresses = mokiutils.CreateRandomAccounts(3)

func TestGetCmdGauges(t *testing.T) {
	desc, _ := GetCmdGauges()
	tcs := map[string]mokicli.QueryCliTestCase[*types.GaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.GaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdToDistributeCoins(t *testing.T) {
	desc, _ := GetCmdToDistributeCoins()
	tcs := map[string]mokicli.QueryCliTestCase[*types.ModuleToDistributeCoinsRequest]{
		"basic test": {
			Cmd: "", ExpectedQuery: &types.ModuleToDistributeCoinsRequest{},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdGaugeByID(t *testing.T) {
	desc, _ := GetCmdGaugeByID()
	tcs := map[string]mokicli.QueryCliTestCase[*types.GaugeByIDRequest]{
		"basic test": {
			Cmd: "1", ExpectedQuery: &types.GaugeByIDRequest{Id: 1},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdActiveGauges(t *testing.T) {
	desc, _ := GetCmdActiveGauges()
	tcs := map[string]mokicli.QueryCliTestCase[*types.ActiveGaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.ActiveGaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdActiveGaugesPerDenom(t *testing.T) {
	desc, _ := GetCmdActiveGaugesPerDenom()
	tcs := map[string]mokicli.QueryCliTestCase[*types.ActiveGaugesPerDenomRequest]{
		"basic test": {
			Cmd: "umoki --offset=2",
			ExpectedQuery: &types.ActiveGaugesPerDenomRequest{
				Denom:      "umoki",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdUpcomingGauges(t *testing.T) {
	desc, _ := GetCmdUpcomingGauges()
	tcs := map[string]mokicli.QueryCliTestCase[*types.UpcomingGaugesRequest]{
		"basic test": {
			Cmd: "--offset=2",
			ExpectedQuery: &types.UpcomingGaugesRequest{
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdUpcomingGaugesPerDenom(t *testing.T) {
	desc, _ := GetCmdUpcomingGaugesPerDenom()
	tcs := map[string]mokicli.QueryCliTestCase[*types.UpcomingGaugesPerDenomRequest]{
		"basic test": {
			Cmd: "umoki --offset=2",
			ExpectedQuery: &types.UpcomingGaugesPerDenomRequest{
				Denom:      "umoki",
				Pagination: &query.PageRequest{Key: []uint8{}, Offset: 2, Limit: 100},
			}},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}
