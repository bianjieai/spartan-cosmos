package keeper

import (
	"context"

	"github.com/bianjieai/iritamod/modules/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.MsgServer = Keeper{}

func (m Keeper) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	id, err := m.Keeper.CreateValidator(ctx, *msg)
	if err != nil {
		return &types.MsgCreateValidatorResponse{}, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgCreateValidatorResponse{
		Id: id.String(),
	}, nil
}

func (m Keeper) UpdateValidator(_ context.Context, _ *types.MsgUpdateValidator) (*types.MsgUpdateValidatorResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to update validator")
}

func (m Keeper) RemoveValidator(_ context.Context, _ *types.MsgRemoveValidator) (*types.MsgRemoveValidatorResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to remove validator")
}

func (m Keeper) GrantNode(_ context.Context, _ *types.MsgGrantNode) (*types.MsgGrantNodeResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to grant node")
}

func (m Keeper) RevokeNode(_ context.Context, _ *types.MsgRevokeNode) (*types.MsgRevokeNodeResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to revoke node")
}
