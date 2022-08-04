package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ govtypes.StakingKeeper = Keeper{}

//TODO
func (m Keeper) IterateBondedValidatorsByPower(
	sdk.Context,
	func(index int64, validator stakingtypes.ValidatorI) (stop bool),
) {
}

//TODO
func (m Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	return sdk.ZeroInt()
}

//TODO
func (m Keeper) IterateDelegations(
	ctx sdk.Context,
	delegator sdk.AccAddress,
	fn func(index int64, delegation stakingtypes.DelegationI) (stop bool),
) {
}
