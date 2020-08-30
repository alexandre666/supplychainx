package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// Append a product
func (k Keeper) AppendProduct(ctx sdk.Context, product types.Product) (alreadyExist bool) {
	// Check product doesn't exist
	_, alreadyExist = k.GetProduct(ctx, product.GetName())
	if alreadyExist {
		return alreadyExist
	}

	// Set the product
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalProduct(k.cdc, product)
	store.Set(types.GetProductKey(product.GetName()), bz)

	return false
}

// Fetch one product from its name
func (k Keeper) GetProduct(ctx sdk.Context, name string) (product types.Product, found bool) {
	store := ctx.KVStore(k.storeKey)

	// Search the value
	value := store.Get(types.GetProductKey(name))
	if value == nil {
		return product, false
	}

	// Return the value
	product = types.MustUnmarshalProduct(k.cdc, value)
	return product, true
}

// Increase the count of a product
func (k Keeper) IncreaseProductCount(ctx sdk.Context, name string) (found bool) {
	store := ctx.KVStore(k.storeKey)

	// Search the value
	value := store.Get(types.GetProductKey(name))
	if value == nil {
		return false
	}

	// Increase the unit count and push back the product
	product := types.MustUnmarshalProduct(k.cdc, value)
	product.IncreaseUnit()
	bz := types.MustMarshalProduct(k.cdc, product)
	store.Set(types.GetProductKey(name), bz)

	return true
}

// Retrieve all products
func (k Keeper) GetAllProducts(ctx sdk.Context) (products []types.Product) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ProductsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		product := types.MustUnmarshalProduct(k.cdc, iterator.Value())
		products = append(products, product)
	}

	return products
}
