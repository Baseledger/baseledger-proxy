package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
	"github.com/unibrightio/baseledger/x/baseledger/types"
)

func (k msgServer) CreateBaseledgerTransaction(goCtx context.Context, msg *types.MsgCreateBaseledgerTransaction) (*types.MsgCreateBaseledgerTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	fmt.Printf("BASELEDGER MODULE ACCOUNT %v %v\n", moduleAcct, moduleAcct.String())

	txCreatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)

	fmt.Printf("TX CREATOR ACCOUNT %v %v\n", txCreatorAddress, txCreatorAddress.String())

	coins := k.bankKeeper.GetAllBalances(ctx, txCreatorAddress)

	fmt.Printf("BALANCES BEFORE SENDING %v\n", coins)

	coins = k.bankKeeper.GetAllBalances(ctx, moduleAcct)

	fmt.Printf("BALANCE MODULE ACCOUNT BEFORE SENDING %v\n", coins)

	if err != nil {
		panic(err)
	}

	coinFee, err := sdk.ParseCoinsNormalized("1token")
	if err != nil {
		panic(err)
	}

	fmt.Printf("SENDING %v", coinFee)
	sdkError := k.bankKeeper.SendCoins(ctx, txCreatorAddress, moduleAcct, coinFee)
	if sdkError != nil {
		fmt.Printf("SEND COINS ERROR %v\n", sdkError.Error())
		return nil, sdkError
	}

	coins = k.bankKeeper.GetAllBalances(ctx, txCreatorAddress)

	fmt.Printf("BALANCES AFTER SENDING 2 %v\n", coins)

	coins = k.bankKeeper.GetAllBalances(ctx, moduleAcct)

	fmt.Printf("BALANCE MODULE ACCOUNT AFTER SENDING %v\n", coins)

	var BaseledgerTransaction = types.BaseledgerTransaction{
		Id:                      msg.Id,
		Creator:                 msg.Creator,
		BaseledgerTransactionId: msg.BaseledgerTransactionId,
		Payload:                 msg.Payload,
	}

	id := k.AppendBaseledgerTransaction(
		ctx,
		BaseledgerTransaction,
	)

	return &types.MsgCreateBaseledgerTransactionResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBaseledgerTransaction(goCtx context.Context, msg *types.MsgUpdateBaseledgerTransaction) (*types.MsgUpdateBaseledgerTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var BaseledgerTransaction = types.BaseledgerTransaction{
		Creator:                 msg.Creator,
		Id:                      msg.Id,
		BaseledgerTransactionId: msg.BaseledgerTransactionId,
		Payload:                 msg.Payload,
	}

	// Checks that the element exists
	if !k.HasBaseledgerTransaction(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBaseledgerTransactionOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBaseledgerTransaction(ctx, BaseledgerTransaction)

	return &types.MsgUpdateBaseledgerTransactionResponse{}, nil
}

func (k msgServer) DeleteBaseledgerTransaction(goCtx context.Context, msg *types.MsgDeleteBaseledgerTransaction) (*types.MsgDeleteBaseledgerTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBaseledgerTransaction(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBaseledgerTransactionOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBaseledgerTransaction(ctx, msg.Id)

	return &types.MsgDeleteBaseledgerTransactionResponse{}, nil
}
