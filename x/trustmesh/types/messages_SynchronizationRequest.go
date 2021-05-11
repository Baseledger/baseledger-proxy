package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSynchronizationRequest{}

func NewMsgCreateSynchronizationRequest(creator string, WorkgroupID string, Recipient string, WorkstepType string, BusinessObjectType string, BaseledgerBusinessObjectID string, BusinessObject string, ReferencedBaseledgerBusinessObjectID string) *MsgCreateSynchronizationRequest {
	return &MsgCreateSynchronizationRequest{
		Creator:                              creator,
		WorkgroupID:                          WorkgroupID,
		Recipient:                            Recipient,
		WorkstepType:                         WorkstepType,
		BusinessObjectType:                   BusinessObjectType,
		BaseledgerBusinessObjectID:           BaseledgerBusinessObjectID,
		BusinessObject:                       BusinessObject,
		ReferencedBaseledgerBusinessObjectID: ReferencedBaseledgerBusinessObjectID,
	}
}

func (msg *MsgCreateSynchronizationRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateSynchronizationRequest) Type() string {
	return "CreateSynchronizationRequest"
}

func (msg *MsgCreateSynchronizationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSynchronizationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSynchronizationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSynchronizationRequest{}

func NewMsgUpdateSynchronizationRequest(creator string, id uint64, WorkgroupID string, Recipient string, WorkstepType string, BusinessObjectType string, BaseledgerBusinessObjectID string, BusinessObject string, ReferencedBaseledgerBusinessObjectID string) *MsgUpdateSynchronizationRequest {
	return &MsgUpdateSynchronizationRequest{
		Id:                                   id,
		Creator:                              creator,
		WorkgroupID:                          WorkgroupID,
		Recipient:                            Recipient,
		WorkstepType:                         WorkstepType,
		BusinessObjectType:                   BusinessObjectType,
		BaseledgerBusinessObjectID:           BaseledgerBusinessObjectID,
		BusinessObject:                       BusinessObject,
		ReferencedBaseledgerBusinessObjectID: ReferencedBaseledgerBusinessObjectID,
	}
}

func (msg *MsgUpdateSynchronizationRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSynchronizationRequest) Type() string {
	return "UpdateSynchronizationRequest"
}

func (msg *MsgUpdateSynchronizationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSynchronizationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSynchronizationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateSynchronizationRequest{}

func NewMsgDeleteSynchronizationRequest(creator string, id uint64) *MsgDeleteSynchronizationRequest {
	return &MsgDeleteSynchronizationRequest{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSynchronizationRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSynchronizationRequest) Type() string {
	return "DeleteSynchronizationRequest"
}

func (msg *MsgDeleteSynchronizationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSynchronizationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSynchronizationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
