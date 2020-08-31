package types

// scx module event types
const (
	EventTypeAppendOrganization    = "append_organization"
	EventTypeRelegateOrganization  = "relegate_organization"
	EventTypeReapproveOrganization = "reapprove_organization"
	EventTypeCreateProduct         = "create_product"
	EventTypeCreateUnit            = "create_unit"
	EventTypeTransferUnit          = "transfer_unit"

	AttributeKeyAuthority           = "authority"
	AttributeKeyOrganizationAddress = "organization_address"
	AttributeKeyOrganizationName    = "organization_name"
	AttributeKeyManufacturer        = "manufacturer"
	AttributeKeyProduct             = "product"
	AttributeKeyReference           = "reference"
	AttributeKeyFrom                = "from"
	AttributeKeyTo                  = "to"

	AttributeValueCategory = ModuleName
)
