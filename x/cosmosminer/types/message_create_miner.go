package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMiner = "create_miner"

var _ sdk.Msg = &MsgCreateMiner{}

func NewMsgCreateMiner(creator string, minerTemplateID uint64) *MsgCreateMiner {
	return &MsgCreateMiner{
		Creator:         creator,
		MinerTemplateID: minerTemplateID,
	}
}

func (msg *MsgCreateMiner) Route() string {
	return RouterKey
}

func (msg *MsgCreateMiner) Type() string {
	return TypeMsgCreateMiner
}

func (msg *MsgCreateMiner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMiner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMiner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
