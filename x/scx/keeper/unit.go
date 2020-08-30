package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// Append an unit
func (k Keeper) AppendUnit(ctx sdk.Context, unit types.Unit) (alreadyExist bool) {
	// TODO

	return false
}

// Fetch a unit from its reference
func (k Keeper) GetUnit(ctx sdk.Context, reference string) (unit types.Unit, found bool) {
	// TODO

	return unit, true
}

// Get the list of all the holder
func (k Keeper) GetUnitTrace(ctx sdk.Context, reference string) (trace []sdk.AccAddress, found bool) {
	// TODO

	return trace, false
}

// Get all the component units of the unit
func (k Keeper) GetUnitComponents(ctx sdk.Context, reference string) (components []types.Unit, found bool) {
	// TODO

	return components, false
}

// Transfer the ownership of an unit
func (k Keeper) TransferUnit(ctx sdk.Context, reference string, newHolder sdk.AccAddress) error {
	// TODO

	// Check if the unit is already component of another unit

	return nil
}

// Set the ComponentOf field
func (k Keeper) SetComponentOf(ctx sdk.Context, reference string, componentOfReference string) error {
	// TODO

	// Check if the unit is already component of another unit

	return nil
}
