package keeper

import (
	"context"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/karthik340/game/x/lottery/types"
)

func (k Keeper) ValidateBet(ctx sdk.Context, msg *types.MsgPlaceBet) error {
	params := k.GetParams(ctx)
	// fee should be 5 token and bet size should be in between 1 and 100
	if !msg.GetFee().IsEqual(sdk.NewCoin("token", sdk.NewInt(int64(params.GetMinFee())))) ||
		msg.GetBet().IsLT(sdk.NewCoin("token", sdk.NewInt(int64(params.GetMinBet())))) ||
		msg.GetBet().IsGTE(sdk.NewCoin("token", sdk.NewInt(int64(params.GetMaxBet())))) {
		return types.ErrBetValidationFailed
	}
	return nil
}

func (k Keeper) substituteBet(
	ctx sdk.Context,
	sender sdk.AccAddress,
	existingUser *types.Bet,
	msg *types.MsgPlaceBet,
) (*types.MsgPlaceBetResponse, error) {
	err := k.bankKeeper.SendCoinsFromModuleToAccount( // send both fee and betSize of previous bet from module to user
		ctx,
		types.ModuleName,
		sender,
		sdk.NewCoins(existingUser.Bet.Add(existingUser.Fee)),
	)
	if err != nil {
		return nil, errors.New("cannot send funds to user")
	}

	totalAmount := msg.Fee.Add(msg.Bet)

	// send betSize+Fee from sender to module account
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(totalAmount))
	if err != nil {
		return nil, types.ErrInsufficientFunds
	}

	// use same txNum as bet exists already
	existingUser.Bet = msg.Bet // update with new bet

	k.SetBetInCurrentRound(
		ctx,
		*existingUser,
	)

	return &types.MsgPlaceBetResponse{}, nil
}

func (k Keeper) AddBet(
	ctx sdk.Context,
	sender sdk.AccAddress,
	msg *types.MsgPlaceBet,
) (*types.MsgPlaceBetResponse, error) {
	totalAmount := msg.Fee.Add(msg.Bet)

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(totalAmount))
	if err != nil {
		return nil, types.ErrInsufficientFunds
	}

	txnCounter, _ := k.GetTxnCounter(ctx)

	var user = types.Bet{
		Sender: msg.Sender,
		Fee:    msg.Fee,
		Bet:    msg.Bet,
		TxNum:  txnCounter.Val,
		Status: false,
	}

	k.SetBetInCurrentRound(
		ctx,
		user,
	)

	// increment tx number
	txnCounter.Val += 1

	k.SetTxnCounter(ctx, txnCounter)

	return &types.MsgPlaceBetResponse{}, nil
}

func (k msgServer) PlaceBet(goCtx context.Context, msg *types.MsgPlaceBet) (*types.MsgPlaceBetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.ValidateBet(ctx, msg)
	if err != nil {
		return nil, err
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Check if user bet already exists for current round
	existingBet, betExists := k.GetBetInCurrentRound(
		ctx,
		msg.Sender,
	)

	if betExists {
		return k.substituteBet(ctx, sender, &existingBet, msg)
	}

	return k.AddBet(ctx, sender, msg)
}
