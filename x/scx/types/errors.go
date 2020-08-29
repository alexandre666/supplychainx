package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidOrganization       = sdkerrors.Register(ModuleName, 1, "invalid organization")
	ErrNotAnAuthority            = sdkerrors.Register(ModuleName, 2, "the sender is not an authority")
	ErrOrganizationAlreadyExists = sdkerrors.Register(ModuleName, 3, "the organization already exists")
	ErrOrganizationNotFound      = sdkerrors.Register(ModuleName, 4, "organization not found")
)
