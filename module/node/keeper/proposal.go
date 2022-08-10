package keeper

import (
	"encoding/hex"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/iritamod/modules/node/types"

	"github.com/bianjieai/spartan-cosmos/module/node"
)

func (k *Keeper) handleValidatorCreateProposal(ctx sdk.Context, p *node.CreateValidatorProposal) error {
	pk, ok := p.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Expecting cryptotypes.PubKey, got %T", pk)
	}

	id := tmbytes.HexBytes(tmhash.Sum(pk.Bytes()))
	if err := k.Keeper.CreateValidator(ctx,
		id,
		p.Name,
		"",
		pk,
		p.Power,
		p.Description,
		p.Operator,
	); err != nil {
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
	id, err := hex.DecodeString(p.Id)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid validator id")
	}

	if err := k.Keeper.UpdateValidator(ctx,
		id,
		p.Name,
		"",
		p.Power,
		p.Description,
		p.Operator); err != nil {
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
	id, err := hex.DecodeString(p.Id)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid validator id")
	}

	if err := k.Keeper.RemoveValidator(ctx, id, p.Operator); err != nil {
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
