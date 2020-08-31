# Events

The scx module emits the following events:

## MsgAppendOrganization

| Type     | Attribute Key | Attribute Value    |
|----------|---------------|--------------------|
| append_organization | module     | scx |
| append_organization | authority     | {authorityValdidatorAddresse} |
| append_organization | organization_address     | {organizationAddress} |
| append_organization | organization_name     | {organizationName} |

## MsgChangeOrganizationApproval

**Relegation**

| Type     | Attribute Key | Attribute Value    |
|----------|---------------|--------------------|
| relegate_organization | module     | scx |
| relegate_organization | authority     | {authorityValdidatorAddresse} |
| relegate_organization | organization_address     | {organizationAddress} |

**Reapproval**

| Type     | Attribute Key | Attribute Value    |
|----------|---------------|--------------------|
| reapprove_organization | module     | scx |
| reapprove_organization | authority     | {authorityValdidatorAddresse} |
| reapprove_organization | organization_address     | {organizationAddress} |

## MsgCreateProduct

| Type     | Attribute Key | Attribute Value    |
|----------|---------------|--------------------|
| create_product | module     | scx |
| create_product | manufacturer     | {organizationAddress} |
| create_product | organization_address     | {productName} |

## MsgCreateUnit

| Type     | Attribute Key | Attribute Value    |
|----------|---------------|--------------------|
| create_unit| module     | scx |
| create_unit | reference     | {unitReference} |
| create_unit | product     | {productName} |

## MsgTransferUnit

| Type     | Attribute Key | Attribute Value    |
|----------|---------------|--------------------|
| transfer_unit| module     | scx |
| transfer_unit | reference     | {unitReference} |
| transfer_unit | from     | {organizationAddress} |
| transfer_unit | to     | {organizationAddress} |