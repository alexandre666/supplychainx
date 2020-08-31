package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidOrganization          = sdkerrors.Register(ModuleName, 1, "invalid organization")
	ErrNotAnAuthority               = sdkerrors.Register(ModuleName, 2, "sender is not an authority")
	ErrOrganizationAlreadyExists    = sdkerrors.Register(ModuleName, 3, "organization already exists")
	ErrOrganizationNotFound         = sdkerrors.Register(ModuleName, 4, "organization not found")
	ErrOrganizationAlreadyApproved  = sdkerrors.Register(ModuleName, 5, "organization already approved")
	ErrOrganizationAlreadyRelegated = sdkerrors.Register(ModuleName, 6, "organization already relegated")
	ErrInvalidProduct               = sdkerrors.Register(ModuleName, 7, "invalid product")
	ErrOrganizationNotApproved      = sdkerrors.Register(ModuleName, 8, "the organization is not approved")
	ErrProductAlreadyExist          = sdkerrors.Register(ModuleName, 9, "the product already exist")
	ErrInvalidUnit                  = sdkerrors.Register(ModuleName, 10, "invalid unit")
	ErrInvalidTransfer              = sdkerrors.Register(ModuleName, 11, "invalid transfer")
	ErrUnitNotFound                 = sdkerrors.Register(ModuleName, 12, "unit not found")
	ErrUnitComponentOfAnotherUnit   = sdkerrors.Register(ModuleName, 13, "unit the unit is component of another unit")
	ErrComponentOfNotFound          = sdkerrors.Register(ModuleName, 14, "the component reference has been not found")
	ErrProductNotFound              = sdkerrors.Register(ModuleName, 15, "product not found")
	ErrComponentNotOwned            = sdkerrors.Register(ModuleName, 16, "component is not owned by the manufacturer")
	ErrUnitNotOwned                 = sdkerrors.Register(ModuleName, 17, "unit is not owned by the holder")
)
