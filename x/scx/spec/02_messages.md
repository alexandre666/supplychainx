# Messages

In this section, we describe the processing of the messages and the corresponding updates to the state.

## MsgAppendOrganization

An organization is appended by an authority (POA validator) in the system with the `MsgAppendOrganization` message.

```go
type MsgAppendOrganization struct {
	Organization Organization
	Authority    sdk.ValAddress
}
```

This message is expected to fail if:

- The authority is not in the validator set of the POA module
- The organization already exists

This message appends the organization in the organizations store.

## MsgChangeOrganizationApproval

The approval status of an organization can be changed by an authority with the `MsgChangeOrganizationApproval` message.
The organization is relegated or reapproved.

```go
type MsgChangeOrganizationApproval struct {
	Organization sdk.AccAddress
	Authority    sdk.ValAddress
	Approve      bool
}

```

This message is expected to fail if:

- The authority is not in the validator set of the POA module
- The state of the organization doesn't change

This message changes the approval state of the organization.

## MsgCreateProduct

A new product template can be registered by an organization with the `MsgCreateProduct` message.

```go
type MsgCreateProduct struct {
	Product Product
}
```

This message is expected to fail if:

- The product name already exists
- The organization doesn't exist or is relegated

The message appends the product in the products store.

## MsgCreateUnit

A manufactured unit of a product can be registered with the `MsgCreateUnit` message.

```go
type MsgCreateUnit struct {
	ProductName  string
	Manufacturer sdk.AccAddress
	Details      string
	Components   []string
}
```

This message is expected to fail if:

- The manufacturer is not the manufacturer of the product
- The product doesn't exist
- The manufacturer is relegated
- A component of the unit doesn't exist
- A component of the unit is not held by the manufacturer
- A component of the unit is already the component of another unit

This message registers the unit in the units store, increases the unit count of the product, and update the components of the unit.

## MsgTransferUnit

The ownership of a unit can be transfer with the `MsgTransferUnit` message.

```go
type MsgTransferUnit struct {
	UnitReference string
	Holder        sdk.AccAddress
	NewHolder     sdk.AccAddress
}
```

This message is expected to fail if:

- The holder is not the actual holder of the unit
- The holder is relegated
- The new holder is not an approved organization
- The unit is component of another unit (in this case it cannot be transferred anymore)

This message updates the holder of the unit and updates the holder history of the unit.
