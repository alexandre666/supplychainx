package types

// scx module event types
const (
	EventTypeAppendOrganization    = "append_organization"
	EventTypeRelegateOrganization  = "relegate_organization"
	EventTypeReapproveOrganization = "reapprove_organization"
	EventTypeCreateProduct         = "create_product"

	AttributeKeyAuthority           = "authority"
	AttributeKeyOrganizationAddress = "organization_address"
	AttributeKeyOrganizationName    = "Organization_name"
	AttributeKeyManufacturer        = "manufacturer"
	AttributeKeyProduct             = "product"

	AttributeValueCategory = ModuleName
)
