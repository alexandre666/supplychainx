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
