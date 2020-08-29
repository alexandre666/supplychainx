package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// Keeper of the scx store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
	poa      types.PoaKeeper
}

// NewKeeper creates a scx keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, poa types.PoaKeeper) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
		poa:      poa,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Verify if an address is an authority (an authority is in the validator set)
func (k Keeper) IsAuthority(ctx sdk.Context, addr sdk.ValAddress) bool {
	_, isInValidatorSet := k.poa.GetValidator(ctx, addr)

	return isInValidatorSet
}
