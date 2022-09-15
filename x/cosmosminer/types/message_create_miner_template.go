package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMinerTemplate = "create_miner_template"

var _ sdk.Msg = &MsgCreateMinerTemplate{}

func NewMsgCreateMinerTemplate(creator string, name string, price uint64, prodPerSecond uint64, minClaimTime uint64) *MsgCreateMinerTemplate {
	return &MsgCreateMinerTemplate{
		Creator:       creator,
		Name:          name,
		Price:         price,
		ProdPerSecond: prodPerSecond,
		MinClaimTime:  minClaimTime,
	}
}

func (msg *MsgCreateMinerTemplate) Route() string {
	return RouterKey
}

func (msg *MsgCreateMinerTemplate) Type() string {
	return TypeMsgCreateMinerTemplate
}

func (msg *MsgCreateMinerTemplate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMinerTemplate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMinerTemplate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
