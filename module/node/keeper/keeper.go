package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	nodekeeper "github.com/bianjieai/iritamod/modules/node/keeper"
	"github.com/bianjieai/iritamod/modules/node/types"
	cautil "github.com/bianjieai/iritamod/utils/ca"
)

type Keeper struct {
	nodekeeper.Keeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, ps paramtypes.Subspace) Keeper {
	return Keeper{Keeper: nodekeeper.NewKeeper(cdc, storeKey, ps)}
}

//VerifyCert override nodekeeper.Keeper.VerifyCert, don't Verify cert used by root cert
func (k *Keeper) VerifyCert(ctx sdk.Context, certStr string) (cert cautil.Cert, err error) {
	cert, err = cautil.ReadCertificateFromMem([]byte(certStr))
	if err != nil {
		return cert, sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
	}
	return cert, nil
}
