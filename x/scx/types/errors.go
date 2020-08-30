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
	ErrOrganizationNotApproved      = sdkerrors.Register(ModuleName, 8, "the orfanization is not approved")
	ErrProductAlreadyExist          = sdkerrors.Register(ModuleName, 9, "the product already exist")
)
