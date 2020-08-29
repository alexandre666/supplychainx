package types

// scx module event types
const (
	EventTypeAppendOrganization    = "append_organization"
	EventTypeRelegateOrganization  = "relegate_organization"
	EventTypeReapproveOrganization = "reapprove_organization"

	AttributeKeyAuthority           = "authority"
	AttributeKeyOrganizationAddress = "organization_address"
	AttributeKeyOrganizationName    = "Organization_name"

	AttributeValueCategory = ModuleName
)
