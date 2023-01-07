package cli

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/petri-labs/mokita/osmoutils"
	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/petri-labs/mokita/x/lockup/types"
)

var testAddresses = osmoutils.CreateRandomAccounts(3)

func TestLockTokensCmd(t *testing.T) {
	desc, _ := NewLockTokensCmd()
	tcs := map[string]mokicli.TxCliTestCase[*types.MsgLockTokens]{
		"lock 201stake tokens for 1 day": {
			Cmd: "201umoki --duration=24h --from=" + testAddresses[0].String(),
			ExpectedMsg: &types.MsgLockTokens{
				Owner:    testAddresses[0].String(),
				Duration: time.Hour * 24,
				Coins:    sdk.NewCoins(sdk.NewInt64Coin("umoki", 201)),
			},
		},
	}
	mokicli.RunTxTestCases(t, desc, tcs)
}

func TestBeginUnlockingAllCmd(t *testing.T) {
	desc, _ := NewBeginUnlockingAllCmd()
	tcs := map[string]mokicli.TxCliTestCase[*types.MsgBeginUnlockingAll]{
		"basic test": {
			Cmd: "--from=" + testAddresses[0].String(),
			ExpectedMsg: &types.MsgBeginUnlockingAll{
				Owner: testAddresses[0].String(),
			},
		},
	}
	mokicli.RunTxTestCases(t, desc, tcs)
}

func TestBeginUnlockingByIDCmd(t *testing.T) {
	desc, _ := NewBeginUnlockByIDCmd()
	tcs := map[string]mokicli.TxCliTestCase[*types.MsgBeginUnlocking]{
		"basic test no coins": {
			Cmd: "10 --from=" + testAddresses[0].String(),
			ExpectedMsg: &types.MsgBeginUnlocking{
				Owner: testAddresses[0].String(),
				ID:    10,
				Coins: sdk.Coins(nil),
			},
		},
		"basic test w/ coins": {
			Cmd: "10 --amount=5umoki --from=" + testAddresses[0].String(),
			ExpectedMsg: &types.MsgBeginUnlocking{
				Owner: testAddresses[0].String(),
				ID:    10,
				Coins: sdk.NewCoins(sdk.NewInt64Coin("umoki", 5)),
			},
		},
	}
	mokicli.RunTxTestCases(t, desc, tcs)
}

func TestModuleBalanceCmd(t *testing.T) {
	desc, _ := GetCmdModuleBalance()
	tcs := map[string]mokicli.QueryCliTestCase[*types.ModuleBalanceRequest]{
		"basic test": {
			Cmd:           "",
			ExpectedQuery: &types.ModuleBalanceRequest{},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestAccountUnlockingCoinsCmd(t *testing.T) {
	desc, _ := GetCmdAccountUnlockingCoins()
	tcs := map[string]mokicli.QueryCliTestCase[*types.AccountUnlockingCoinsRequest]{
		"basic test": {
			Cmd: testAddresses[0].String(),
			ExpectedQuery: &types.AccountUnlockingCoinsRequest{
				Owner: testAddresses[0].String(),
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestCmdAccountLockedPastTime(t *testing.T) {
	desc, _ := GetCmdAccountLockedPastTime()
	tcs := map[string]mokicli.QueryCliTestCase[*types.AccountLockedPastTimeRequest]{
		"basic test": {
			Cmd: testAddresses[0].String() + " 1670431012",
			ExpectedQuery: &types.AccountLockedPastTimeRequest{
				Owner:     testAddresses[0].String(),
				Timestamp: time.Unix(1670431012, 0),
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestCmdAccountLockedPastTimeNotUnlockingOnly(t *testing.T) {
	desc, _ := GetCmdAccountLockedPastTimeNotUnlockingOnly()
	tcs := map[string]mokicli.QueryCliTestCase[*types.AccountLockedPastTimeNotUnlockingOnlyRequest]{
		"basic test": {
			Cmd: testAddresses[0].String() + " 1670431012",
			ExpectedQuery: &types.AccountLockedPastTimeNotUnlockingOnlyRequest{
				Owner:     testAddresses[0].String(),
				Timestamp: time.Unix(1670431012, 0),
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}

func TestCmdTotalLockedByDenom(t *testing.T) {
	desc, _ := GetCmdTotalLockedByDenom()
	tcs := map[string]mokicli.QueryCliTestCase[*types.LockedDenomRequest]{
		"basic test": {
			Cmd: "umoki --min-duration=1s",
			ExpectedQuery: &types.LockedDenomRequest{
				Denom:    "umoki",
				Duration: time.Second,
			},
		},
	}
	mokicli.RunQueryTestCases(t, desc, tcs)
}
