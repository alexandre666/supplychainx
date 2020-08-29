package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidOrganization = sdkerrors.Register(ModuleName, 1, "invalid organization")
)
