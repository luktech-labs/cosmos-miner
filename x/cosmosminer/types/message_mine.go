package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMine = "mine"

var _ sdk.Msg = &MsgMine{}

func NewMsgMine(creator string, minerID uint64) *MsgMine {
	return &MsgMine{
		Creator: creator,
		MinerID: minerID,
	}
}

func (msg *MsgMine) Route() string {
	return RouterKey
}

func (msg *MsgMine) Type() string {
	return TypeMsgMine
}

func (msg *MsgMine) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMine) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMine) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
