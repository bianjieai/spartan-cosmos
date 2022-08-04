package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	nodekeeper "github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"

	"github.com/bianjieai/spartan-cosmos/module/node"
)

type Keeper struct {
	nodekeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, ps paramtypes.Subspace) Keeper {
	k := nodekeeper.NewKeeper(cdc, storeKey, ps)
	k = k.SetVerifyCertFn(func(ctx sdk.Context, certStr string) (cert cautil.Cert, err error) {
		cert, err = cautil.ReadCertificateFromMem([]byte(certStr))
		if err != nil {
			return cert, sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
		}
		return cert, nil
	})
	return Keeper{Keeper: k}
}

func (k *Keeper) ProposalHandler() govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *node.CreateValidatorProposal:
			return k.handleValidatorCreateProposal(ctx, c)
		case *node.UpdateValidatorProposal:
			return k.handleValidatorUpdateProposal(ctx, c)
		case *node.RemoveValidatorProposal:
			return k.handleValidatorRemoveProposal(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
