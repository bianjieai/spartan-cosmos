package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/spartan-cosmos/module/node"
)

func (k *Keeper) handleValidatorCreateProposal(ctx sdk.Context, p *node.CreateValidatorProposal) error {
	return nil
}

func (k *Keeper) handleValidatorUpdateProposal(ctx sdk.Context, p *node.UpdateValidatorProposal) error {
	return nil
}

func (k *Keeper) handleValidatorRemoveProposal(ctx sdk.Context, p *node.RemoveValidatorProposal) error {
	return nil
}
