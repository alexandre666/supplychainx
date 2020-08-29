package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// Get on organization
func (k Keeper) GetOrganization(ctx sdk.Context, addr sdk.AccAddress) (organization types.Organization, found bool) {
	store := ctx.KVStore(k.storeKey)

	// Search the value
	value := store.Get(types.GetOrganizationKey(addr))
	if value == nil {
		return organization, false
	}

	// Return the value
	organization = types.MustUnmarshalOrganization(k.cdc, value)
	return organization, true
}

// Set organization details
func (k Keeper) SetOrganization(ctx sdk.Context, organization types.Organization) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalOrganization(k.cdc, organization)
	store.Set(types.GetOrganizationKey(organization.GetAddress()), bz)
}

// Get the set of all organizations
func (k Keeper) GetAllOrganizations(ctx sdk.Context) (organizations []types.Organization) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.OrganizationsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		organization := types.MustUnmarshalOrganization(k.cdc, iterator.Value())
		organizations = append(organizations, organization)
	}

	return organizations
}
