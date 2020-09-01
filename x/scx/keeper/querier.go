package keeper

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ltacker/supplychainx/x/scx/types"
)

// NewQuerier creates a new querier for scx clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {

		case types.QueryOrganizations:
			return queryOrganizations(ctx, k)

		case types.QueryOrganization:
			return queryOrganization(ctx, req, k)

		case types.QueryProduct:
			return queryProduct(ctx, req, k)

		case types.QueryProductUnits:
			return queryProductUnits(ctx, req, k)

		case types.QueryUnit:
			return queryUnit(ctx, req, k)

		case types.QueryUnitTrace:
			return queryUnitTrace(ctx, req, k)

		case types.QueryUnitComponents:
			return queryUnitComponents(ctx, req, k)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown scx query endpoint")
		}
	}
}

func queryOrganizations(ctx sdk.Context, k Keeper) ([]byte, error) {
	organizations := k.GetAllOrganizations(ctx)

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, organizations)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryOrganization(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryOrganizationParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	organization, found := k.GetOrganization(ctx, params.OrganizationAddr)
	if !found {
		return nil, types.ErrOrganizationNotFound
	}

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, organization)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryProduct(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryProductParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	product, found := k.GetProduct(ctx, params.ProductName)
	if !found {
		return nil, types.ErrProductNotFound
	}

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, product)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryProductUnits(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryProductParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the product count
	product, found := k.GetProduct(ctx, params.ProductName)
	if !found {
		return nil, types.ErrProductNotFound
	}
	productCount := product.GetUnitCount()

	var productUnits []types.Unit

	// Compute unit reference
	var i uint
	for i = 0; i < uint(productCount); i++ {
		unitRef, err := types.GetUnitReferenceFromProductAndUnitNumber(params.ProductName, i)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error computing a unit reference: %v, %v, %v", params.ProductName, i, err))
		}

		// Retrieve the unit
		unit, found := k.GetUnit(ctx, unitRef)
		if !found {
			panic(fmt.Sprintf("The unit reference %v is referenced but the unit doesn't exist", unitRef))
		}
		productUnits = append(productUnits, unit)
	}

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, productUnits)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryUnit(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryUnitParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Retrieve the unit
	unit, found := k.GetUnit(ctx, params.UnitReference)
	if !found {
		return nil, types.ErrUnitNotFound
	}

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, unit)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryUnitTrace(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryUnitParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the unit trace
	trace, found := k.GetUnitTrace(ctx, params.UnitReference)
	if !found {
		return nil, types.ErrUnitNotFound
	}

	var organizations []types.Organization

	// Retrieve the organization for each address
	for _, orgAddr := range trace {
		// Get the organization
		org, found := k.GetOrganization(ctx, orgAddr)
		if !found {
			panic(fmt.Sprintf("The organization of address %v is referenced in a unit but doesn't exist", orgAddr))
		}

		organizations = append(organizations, org)
	}

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, organizations)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryUnitComponents(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryUnitParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get all components
	comps, found := k.GetUnitComponents(ctx, params.UnitReference)
	if !found {
		return nil, types.ErrUnitNotFound
	}

	// Encode response
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, comps)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
