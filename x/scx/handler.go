package scx

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ltacker/supplychainx/x/scx/keeper"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// NewHandler creates an sdk.Handler for all the scx type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgAppendOrganization:
			return handleMsgAppendOrganization(ctx, k, msg)
		case types.MsgChangeOrganizationApproval:
			return handleMsgChangeOrganizationApproval(ctx, k, msg)
		case types.MsgCreateProduct:
			return handleMsgCreateProduct(ctx, k, msg)
		case types.MsgCreateUnit:
			return handleMsgCreateUnit(ctx, k, msg)
		case types.MsgTransferUnit:
			return handleMsgTransferUnit(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgAppendOrganization(ctx sdk.Context, k keeper.Keeper, msg types.MsgAppendOrganization) (*sdk.Result, error) {
	// Check if the authority is valid
	if !k.IsAuthority(ctx, msg.Authority) {
		return nil, types.ErrNotAnAuthority
	}

	// Check if the organization exist
	_, found := k.GetOrganization(ctx, msg.Organization.GetAddress())
	if found {
		return nil, types.ErrOrganizationAlreadyExists
	}

	// Set the organization
	k.SetOrganization(ctx, msg.Organization)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAppendOrganization,
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority.String()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyOrganizationAddress, msg.Organization.GetAddress().String()),
			sdk.NewAttribute(types.AttributeKeyOrganizationName, msg.Organization.GetName()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgChangeOrganizationApproval(ctx sdk.Context, k keeper.Keeper, msg types.MsgChangeOrganizationApproval) (*sdk.Result, error) {
	// Check if the authority is valid
	if !k.IsAuthority(ctx, msg.Authority) {
		return nil, types.ErrNotAnAuthority
	}

	// Check if the organization exist
	organization, found := k.GetOrganization(ctx, msg.Organization)
	if !found {
		return nil, types.ErrOrganizationNotFound
	}

	// Check if the organization is relegated or reapproved
	if msg.Approve {
		return handleReapproveOrganization(ctx, k, msg, organization)
	} else {
		return handleRelegateOrganization(ctx, k, msg, organization)
	}
}

func handleRelegateOrganization(ctx sdk.Context, k keeper.Keeper, msg types.MsgChangeOrganizationApproval, organization types.Organization) (*sdk.Result, error) {
	// Check if the organization isalready relegated
	if !organization.IsApproved() {
		return nil, types.ErrOrganizationAlreadyRelegated
	}

	// Relegate and update organization
	organization.Relegate()
	k.SetOrganization(ctx, organization)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRelegateOrganization,
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority.String()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyOrganizationAddress, organization.GetAddress().String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleReapproveOrganization(ctx sdk.Context, k keeper.Keeper, msg types.MsgChangeOrganizationApproval, organization types.Organization) (*sdk.Result, error) {
	// Check if the organization isalready approved
	if organization.IsApproved() {
		return nil, types.ErrOrganizationAlreadyApproved
	}

	// Reapprove and update organization
	organization.Approve()
	k.SetOrganization(ctx, organization)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeReapproveOrganization,
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority.String()),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyOrganizationAddress, organization.GetAddress().String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateProduct(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateProduct) (*sdk.Result, error) {
	// The manufacturer must be approved
	organization, found := k.GetOrganization(ctx, msg.Product.GetManufacturer())
	if !found {
		return nil, types.ErrOrganizationNotFound
	}
	if !organization.IsApproved() {
		return nil, types.ErrOrganizationNotApproved
	}

	// Append the product
	exist := k.AppendProduct(ctx, msg.Product)
	if exist {
		return nil, types.ErrProductAlreadyExist
	}

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateProduct,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyManufacturer, msg.Product.Manufacturer.String()),
			sdk.NewAttribute(types.AttributeKeyProduct, msg.Product.GetName()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateUnit(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateUnit) (*sdk.Result, error) {
	// The manufacturer must be approved
	organization, found := k.GetOrganization(ctx, msg.Manufacturer)
	if !found {
		return nil, types.ErrOrganizationNotFound
	}
	if !organization.IsApproved() {
		return nil, types.ErrOrganizationNotApproved
	}

	// Check the product exists
	product, found := k.GetProduct(ctx, msg.ProductName)
	if !found {
		return nil, types.ErrProductNotFound
	}

	// Check the organization is the manufacturer of the product
	if !product.GetManufacturer().Equals(msg.Manufacturer) {
		return nil, types.ErrInvalidOrganization
	}

	// Compute the reference of the product
	unitNumber := product.GetUnitCount()
	reference, err := types.GetUnitReferenceFromProductAndUnitNumber(msg.ProductName, uint(unitNumber))
	if err != nil {
		return nil, types.ErrInvalidUnit
	}

	// Check all components exist, the manufacturer own them and they are not already "component of"
	for _, compRef := range msg.Components {
		// Check the component exists
		component, found := k.GetUnit(ctx, compRef)
		if !found {
			return nil, types.ErrUnitNotFound
		}

		// Check the manufacturer own the unit
		if !component.GetCurrentHolder().Equals(msg.Manufacturer) {
			return nil, types.ErrComponentNotOwned
		}

		if component.IsComponentOf() {
			return nil, types.ErrUnitComponentOfAnotherUnit
		}

		// The the "component of" field of the component
		component.ComponentOf = reference
		k.SetUnit(ctx, component)
	}

	// Create the unit
	unit := types.NewUnit(reference, product, msg.Details, msg.Components)
	alreadyExist := k.AppendUnit(ctx, unit)
	if alreadyExist {
		// Panic if the unit already exists (this should never be the case since the reference is computed for the unit number)
		panic(fmt.Sprintf("The unit %v already exists", reference))
	}

	// Increment the count of unit of the product
	k.IncreaseProductCount(ctx, msg.ProductName)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateUnit,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyReference, reference),
			sdk.NewAttribute(types.AttributeKeyProduct, msg.ProductName),
		),
	})

	// Return the reference
	return &sdk.Result{
		Data:   types.ModuleCdc.MustMarshalBinaryLengthPrefixed(reference),
		Events: ctx.EventManager().Events(),
	}, nil
}

func handleMsgTransferUnit(ctx sdk.Context, k keeper.Keeper, msg types.MsgTransferUnit) (*sdk.Result, error) {
	// Check the holder exists and is approved
	holder, found := k.GetOrganization(ctx, msg.Holder)
	if !found {
		return nil, types.ErrOrganizationNotFound
	}
	if !holder.IsApproved() {
		return nil, types.ErrOrganizationNotApproved
	}

	// Check the new holder exists and is approved
	newHolder, found := k.GetOrganization(ctx, msg.NewHolder)
	if !found {
		return nil, types.ErrOrganizationNotFound
	}
	if !newHolder.IsApproved() {
		return nil, types.ErrOrganizationNotApproved
	}

	// Check the unit exists, the holder owns it and it is not "component of"
	unit, found := k.GetUnit(ctx, msg.UnitReference)
	if !found {
		return nil, types.ErrUnitNotFound
	}
	if !unit.GetCurrentHolder().Equals(msg.Holder) {
		return nil, types.ErrUnitNotOwned
	}
	if unit.IsComponentOf() {
		// If the unit is component of another unit, it cannot be transfered anymore
		return nil, types.ErrUnitComponentOfAnotherUnit
	}

	// Update new hodler
	unit.ChangeHolder(msg.NewHolder)
	k.SetUnit(ctx, unit)

	// Emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferUnit,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyReference, msg.UnitReference),
			sdk.NewAttribute(types.AttributeKeyFrom, msg.Holder.String()),
			sdk.NewAttribute(types.AttributeKeyTo, msg.NewHolder.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
