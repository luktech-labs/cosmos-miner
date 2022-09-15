package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStoredMiner = "create_stored_miner"
	TypeMsgUpdateStoredMiner = "update_stored_miner"
	TypeMsgDeleteStoredMiner = "delete_stored_miner"
)

var _ sdk.Msg = &MsgCreateStoredMiner{}

func NewMsgCreateStoredMiner(
	creator string,
	index string,
	lastClaimed uint64,

) *MsgCreateStoredMiner {
	return &MsgCreateStoredMiner{
		Creator:     creator,
		Index:       index,
		LastClaimed: lastClaimed,
	}
}

func (msg *MsgCreateStoredMiner) Route() string {
	return RouterKey
}

func (msg *MsgCreateStoredMiner) Type() string {
	return TypeMsgCreateStoredMiner
}

func (msg *MsgCreateStoredMiner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStoredMiner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStoredMiner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStoredMiner{}

func NewMsgUpdateStoredMiner(
	creator string,
	index string,
	lastClaimed uint64,

) *MsgUpdateStoredMiner {
	return &MsgUpdateStoredMiner{
		Creator:     creator,
		Index:       index,
		LastClaimed: lastClaimed,
	}
}

func (msg *MsgUpdateStoredMiner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStoredMiner) Type() string {
	return TypeMsgUpdateStoredMiner
}

func (msg *MsgUpdateStoredMiner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStoredMiner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStoredMiner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStoredMiner{}

func NewMsgDeleteStoredMiner(
	creator string,
	index string,

) *MsgDeleteStoredMiner {
	return &MsgDeleteStoredMiner{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteStoredMiner) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStoredMiner) Type() string {
	return TypeMsgDeleteStoredMiner
}

func (msg *MsgDeleteStoredMiner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStoredMiner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStoredMiner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
