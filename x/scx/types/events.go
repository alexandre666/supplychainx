package types

// scx module event types
const (
	EventTypeAppendOrganization    = "append_organization"
	EventTypeRelegateOrganization  = "relegate_organization"
	EventTypeReapproveOrganization = "reapprove_organization"
	EventTypeCreateProduct         = "create_product"
	EventTypeCreateUnit            = "create_unit"

	AttributeKeyAuthority           = "authority"
	AttributeKeyOrganizationAddress = "organization_address"
	AttributeKeyOrganizationName    = "Organization_name"
	AttributeKeyManufacturer        = "manufacturer"
	AttributeKeyProduct             = "product"
	AttributeKeyReference           = "reference"

	AttributeValueCategory = ModuleName
)
