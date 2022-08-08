package keeper

import (
	"context"

	"github.com/bianjieai/iritamod/modules/node/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.MsgServer = Keeper{}

func (m Keeper) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "not supported to create validator")
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
