package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
	"github.com/bianjieai/spartan-cosmos/module/node"
)

func (k *Keeper) handleValidatorCreateProposal(ctx sdk.Context, p *node.CreateValidatorProposal) error {
	id, err := k.Keeper.CreateValidator(ctx, p.ToMsgCreateValidator())
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, id.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, p.Operator),
		),
	})
	return nil
}

func (k *Keeper) handleValidatorUpdateProposal(ctx sdk.Context, p *node.UpdateValidatorProposal) error {
	err := k.Keeper.UpdateValidator(ctx, p.ToMsgUpdateValidator())
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, p.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, p.Operator),
		),
	})
	return nil
}

func (k *Keeper) handleValidatorRemoveProposal(ctx sdk.Context, p *node.RemoveValidatorProposal) error {
	err := k.Keeper.RemoveValidator(ctx, p.ToMsgRemoveValidator())
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, p.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, p.Operator),
		),
	})
	return nil
}
