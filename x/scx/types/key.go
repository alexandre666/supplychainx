package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "scx"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

var (
	// Prefix for each key to an organization
	OrganizationsKey = []byte{0x21}
)

// Get the key for the organization with address
func GetOrganizationKey(orgAddr sdk.AccAddress) []byte {
	return append(OrganizationsKey, orgAddr.Bytes()...)
}
