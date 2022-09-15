package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/luktech-labs/cosmos-miner/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMinerTemplate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMinerTemplate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMinerTemplate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMinerTemplate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
