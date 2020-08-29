package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	poatypes "github.com/ltacker/poa/types"
)

//We use the PoaKeeper to verify if a regulator is a validator in the system
type PoaKeeper interface {
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator poatypes.Validator, found bool)
}
