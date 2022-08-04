package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/node/types"
	"github.com/bianjieai/spartan-cosmos/module/node"
)

func (k *Keeper) handleValidatorCreateProposal(ctx sdk.Context, p *node.CreateValidatorProposal) error {
	msg := &types.MsgCreateValidator{
		Name:        p.Name,
		Certificate: p.Certificate,
		Power:       p.Power,
		Description: p.Description,
		Operator:    p.Operator,
	}

	id, err := k.Keeper.CreateValidator(ctx, *msg)
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
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return nil
}

func (k *Keeper) handleValidatorUpdateProposal(ctx sdk.Context, p *node.UpdateValidatorProposal) error {
	msg := &types.MsgUpdateValidator{
		Id:          p.Id,
		Name:        p.Name,
		Certificate: p.Certificate,
		Power:       p.Power,
		Description: p.Description,
		Operator:    p.Operator,
	}

	err := k.Keeper.UpdateValidator(ctx, *msg)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return nil
}

func (k *Keeper) handleValidatorRemoveProposal(ctx sdk.Context, p *node.RemoveValidatorProposal) error {
	msg := &types.MsgRemoveValidator{
		Id: p.Id,
	}

	err := k.Keeper.RemoveValidator(ctx, *msg)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveValidator,
			sdk.NewAttribute(types.AttributeKeyValidator, msg.Id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})
	return nil
}
