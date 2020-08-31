<!--
order: 1
-->

# State

## Organization

Organizations are the stakeholders that can create, manufacture, and transfer products.
There are appended by authorities and are immutable. Once created, an organization cannot be deleted.
If the organization misbehaves, it can be relegated by an authority. Once relegated, an organization no longer can create, manufacture, transfer, or get transferred a product.

- Organizations: `0x21 | OrganizationAddr -> amino(organization)`

```go
type Organization struct {
	Address     sdk.AccAddress
	Name        string
	Description string
	Approved    bool
}
```

## Product

The product data describes a product manufactured by an organization.
For simplicity, the product is identified by its name. There is only one product per name and one organization per product.
The product is immutable.

- Products: `0x22 | ProductName -> amino(product)`

```go
type Product struct {
	Name         string
	Description  string
	Manufacturer sdk.AccAddress
	UnitCount    uint64
} 
```

## Unit

The unit data describes a manufactured unit of a specific product. A unit is identified by a reference. It can only be created by the product manufactured and it can be transferred.
A unit can be composed of other units and can be the component of another unit. Once the unit is a component of another unit, it cannot be transferred anymore.

The reference is computed from the following hash:
`first16Bytes(sha256(productName,unitNumber))`

- Units: `0x23 | UnitReference -> amino(unit)`

```go
type Unit struct {
	Reference     string
	Product       string
	Details       string
	Components    []string
	Holder        sdk.AccAddress
	HolderHistory []sdk.AccAddress
	ComponentOf   string
}
```