package model_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	appParams "github.com/petri-labs/mokita/app/params"
	clmodel "github.com/petri-labs/mokita/x/concentrated-liquidity/model"
	"github.com/petri-labs/mokita/x/concentrated-liquidity/types"
)

func TestMsgCreateConcentratedPool(t *testing.T) {
	appParams.SetAddressPrefixes()
	pk1 := ed25519.GenPrivKey().PubKey()
	addr1 := sdk.AccAddress(pk1.Address()).String()
	invalidAddr := sdk.AccAddress("invalid")

	tests := []struct {
		name       string
		msg        clmodel.MsgCreateConcentratedPool
		expectPass bool
	}{
		{
			name: "proper msg",
			msg: clmodel.MsgCreateConcentratedPool{
				Sender:      addr1,
				Denom0:      ETH,
				Denom1:      USDC,
				TickSpacing: DefaultTickSpacing,
			},
			expectPass: true,
		},
		{
			name: "invalid sender",
			msg: clmodel.MsgCreateConcentratedPool{
				Sender:      invalidAddr.String(),
				Denom0:      ETH,
				Denom1:      USDC,
				TickSpacing: DefaultTickSpacing,
			},
			expectPass: false,
		},
		{
			name: "missing denom1",
			msg: clmodel.MsgCreateConcentratedPool{
				Sender: invalidAddr.String(),
				Denom0: ETH,
			},
			expectPass: false,
		},
		{
			name: "missing denom0",
			msg: clmodel.MsgCreateConcentratedPool{
				Sender: invalidAddr.String(),
				Denom1: USDC,
			},
			expectPass: false,
		},
		{
			name: "missing sender",
			msg: clmodel.MsgCreateConcentratedPool{
				Denom0: ETH,
				Denom1: USDC,
			},
			expectPass: false,
		},
	}

	for _, test := range tests {
		msg := test.msg

		if test.expectPass {
			require.NoError(t, test.msg.ValidateBasic(), "test: %v", test.name)
			require.Equal(t, msg.Route(), types.RouterKey)
			require.Equal(t, msg.Type(), "create_concentrated_pool")
			signers := msg.GetSigners()
			require.Equal(t, len(signers), 1)
			require.Equal(t, signers[0].String(), addr1)
		} else {
			require.Error(t, test.msg.ValidateBasic(), "test: %v", test.name)
		}
	}
}
