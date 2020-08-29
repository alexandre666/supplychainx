package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Query endpoints supported by the scx querier
const (
	QueryOrganizations = "organizations"
	QueryOrganization  = "organization"
)

// Defines the params for the following queries:
// - 'custom/scx/organization'
type QueryOrganizationParams struct {
	OrganizationAddr sdk.AccAddress
}

func NewQueryOrganizationParams(organizationAddr sdk.AccAddress) QueryOrganizationParams {
	return QueryOrganizationParams{
		OrganizationAddr: organizationAddr,
	}
}
