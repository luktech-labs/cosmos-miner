package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/luktech-labs/cosmos-miner/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateStoredMiner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateStoredMiner
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateStoredMiner{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateStoredMiner{
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

func TestMsgUpdateStoredMiner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateStoredMiner
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateStoredMiner{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateStoredMiner{
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

func TestMsgDeleteStoredMiner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteStoredMiner
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteStoredMiner{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteStoredMiner{
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
