package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// Append an unit
func (k Keeper) AppendUnit(ctx sdk.Context, unit types.Unit) (alreadyExist bool) {
	// Check unit doesn't exist
	_, alreadyExist = k.GetUnit(ctx, unit.GetReference())
	if alreadyExist {
		return alreadyExist
	}

	// Set the unit
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalUnit(k.cdc, unit)
	store.Set(types.GetUnitKeyFromReference(unit.GetReference()), bz)

	return false
}

// Set unit
func (k Keeper) SetUnit(ctx sdk.Context, unit types.Unit) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalUnit(k.cdc, unit)
	store.Set(types.GetUnitKeyFromReference(unit.GetReference()), bz)
}

// Fetch a unit from its reference
func (k Keeper) GetUnit(ctx sdk.Context, reference string) (unit types.Unit, found bool) {
	store := ctx.KVStore(k.storeKey)

	// Search the value
	value := store.Get(types.GetUnitKeyFromReference(reference))
	if value == nil {
		return unit, false
	}

	// Return the value
	unit = types.MustUnmarshalUnit(k.cdc, value)
	return unit, true
}

// Get the list of all the holder
func (k Keeper) GetUnitTrace(ctx sdk.Context, reference string) (trace []sdk.AccAddress, found bool) {
	// Check unit exists
	unit, found := k.GetUnit(ctx, reference)
	if !found {
		return trace, found
	}

	// Get the history
	trace = unit.GetHolderHistrory()

	// Append the current holder to the trace
	trace = append(trace, unit.GetCurrentHolder())

	return trace, true
}

// Get all the component units of the unit
func (k Keeper) GetUnitComponents(ctx sdk.Context, reference string) (components []types.Unit, found bool) {
	// Check unit exists
	unit, found := k.GetUnit(ctx, reference)
	if !found {
		return components, found
	}

	// Get the component references
	componentReferences := unit.GetComponents()

	// Get all the components
	for _, compRef := range componentReferences {
		comp, found := k.GetUnit(ctx, compRef)
		if !found {
			panic("An unit contains a non existing component")
		}
		components = append(components, comp)
	}

	// Get all the sub components
	for _, compRef := range componentReferences {
		subComps, found := k.GetUnitComponents(ctx, compRef)
		if !found {
			panic("An unit contains a non existing component")
		}
		components = append(components, subComps...)
	}

	return components, true
}
