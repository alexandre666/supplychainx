package types

import (
	"fmt"

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

	// Products prefix
	ProductsKey = []byte{0x22}

	// Units prefix
	UnitsKey = []byte{0x22}
)

// Get the key for the organization with address
func GetOrganizationKey(orgAddr sdk.AccAddress) []byte {
	return append(OrganizationsKey, orgAddr.Bytes()...)
}

// Get the key for the products with the product name
func GetProductKey(productName string) []byte {
	return append(ProductsKey, []byte(productName)...)
}

// Get the key for the unit from the product name and the unit number
func GetUnitKeyFromProductAndUnitNumber(productName string, unitNumber uint) []byte {
	reference, err := GetUnitReferenceFromProductAndUnitNumber(productName, unitNumber)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error decoding unit reference: %v", err))
	}

	return append(UnitsKey, []byte(reference)...)
}

// Get the key for the unit from its reference
func GetUnitKeyFromReference(reference string) []byte {
	return append(UnitsKey, []byte(reference)...)
}
