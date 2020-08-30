package keeper

import (
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

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, organization)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryProduct(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	// TODO
	var params types.QueryProductParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	return nil, nil
}

func queryProductUnits(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	// TODO
	var params types.QueryProductParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the product count

	// Compute unit reference

	// Retrieve unit

	return nil, nil
}

func queryUnit(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	// TODO
	var params types.QueryUnitParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get unit

	return nil, nil
}

func queryUnitTrace(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	// TODO
	var params types.QueryUnitParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get all the holder

	// Retrieve the organization from the address

	return nil, nil
}

func queryUnitComponents(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	// TODO
	var params types.QueryUnitParams

	// Unmarschal request
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get all components

	return nil, nil
}
