package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Query endpoints supported by the scx querier
const (
	QueryOrganizations  = "organizations"
	QueryOrganization   = "organization"
	QueryProduct        = "product"
	QueryProductUnits   = "product-units"
	QueryUnit           = "unit"
	QueryUnitTrace      = "unit-trace"
	QueryUnitComponents = "unit-components"
)

// Defines the params for the following querie:
// - 'custom/scx/organization'
type QueryOrganizationParams struct {
	OrganizationAddr sdk.AccAddress
}

func NewQueryOrganizationParams(organizationAddr sdk.AccAddress) QueryOrganizationParams {
	return QueryOrganizationParams{
		OrganizationAddr: organizationAddr,
	}
}

// Defines the params for the following querie:
// - 'custom/scx/product'
// - 'custom/scx/product-units'
type QueryProductParams struct {
	ProductName string
}

func NewQueryProductParams(productName string) QueryProductParams {
	return QueryProductParams{
		ProductName: productName,
	}
}

// Defines the params for the following querie:
// - 'custom/scx/unit'
// - 'custom/scx/unit-trace'
// - 'custom/scx/unit-components'
type QueryUnitParams struct {
	UnitReference string
}

func NewQueryUnitParams(unitReference string) QueryUnitParams {
	return QueryUnitParams{
		UnitReference: unitReference,
	}
}
