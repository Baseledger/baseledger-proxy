package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSynchronizationFeedback{}

func NewMsgCreateSynchronizationFeedback(creator string, Approved string, BusinessObject string, BaseledgerBusinessObjectIDOfApprovedObject string, Workgroup string, Recipient string, HashOfObjectToApprove string, OriginalBaseledgerTransactionID string, OriginalOffchainProcessMessageID string, FeedbackMessage string) *MsgCreateSynchronizationFeedback {
	return &MsgCreateSynchronizationFeedback{
		Creator:        creator,
		Approved:       Approved,
		BusinessObject: BusinessObject,
		BaseledgerBusinessObjectIDOfApprovedObject: BaseledgerBusinessObjectIDOfApprovedObject,
		Workgroup:                        Workgroup,
		Recipient:                        Recipient,
		HashOfObjectToApprove:            HashOfObjectToApprove,
		OriginalBaseledgerTransactionID:  OriginalBaseledgerTransactionID,
		OriginalOffchainProcessMessageID: OriginalOffchainProcessMessageID,
		FeedbackMessage:                  FeedbackMessage,
	}
}

func (msg *MsgCreateSynchronizationFeedback) Route() string {
	return RouterKey
}

func (msg *MsgCreateSynchronizationFeedback) Type() string {
	return "CreateSynchronizationFeedback"
}

func (msg *MsgCreateSynchronizationFeedback) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSynchronizationFeedback) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSynchronizationFeedback) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSynchronizationFeedback{}

func NewMsgUpdateSynchronizationFeedback(creator string, id uint64, Approved string, BusinessObject string, BaseledgerBusinessObjectIDOfApprovedObject string, Workgroup string, Recipient string, HashOfObjectToApprove string, OriginalBaseledgerTransactionID string, OriginalOffchainProcessMessageID string, FeedbackMessage string) *MsgUpdateSynchronizationFeedback {
	return &MsgUpdateSynchronizationFeedback{
		Id:             id,
		Creator:        creator,
		Approved:       Approved,
		BusinessObject: BusinessObject,
		BaseledgerBusinessObjectIDOfApprovedObject: BaseledgerBusinessObjectIDOfApprovedObject,
		Workgroup:                        Workgroup,
		Recipient:                        Recipient,
		HashOfObjectToApprove:            HashOfObjectToApprove,
		OriginalBaseledgerTransactionID:  OriginalBaseledgerTransactionID,
		OriginalOffchainProcessMessageID: OriginalOffchainProcessMessageID,
		FeedbackMessage:                  FeedbackMessage,
	}
}

func (msg *MsgUpdateSynchronizationFeedback) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSynchronizationFeedback) Type() string {
	return "UpdateSynchronizationFeedback"
}

func (msg *MsgUpdateSynchronizationFeedback) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSynchronizationFeedback) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSynchronizationFeedback) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateSynchronizationFeedback{}

func NewMsgDeleteSynchronizationFeedback(creator string, id uint64) *MsgDeleteSynchronizationFeedback {
	return &MsgDeleteSynchronizationFeedback{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteSynchronizationFeedback) Route() string {
	return RouterKey
}

func (msg *MsgDeleteSynchronizationFeedback) Type() string {
	return "DeleteSynchronizationFeedback"
}

func (msg *MsgDeleteSynchronizationFeedback) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteSynchronizationFeedback) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteSynchronizationFeedback) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
