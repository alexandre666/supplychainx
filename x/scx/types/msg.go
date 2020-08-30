package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var _ sdk.Msg = &MsgAppendOrganization{}
var _ sdk.Msg = &MsgChangeOrganizationApproval{}

/**
 * MsgAppendOrganization
 */
type MsgAppendOrganization struct {
	Organization Organization   `json:"organization"`
	Authority    sdk.ValAddress `json:"authority"`
}

func NewMsgAppendOrganization(organization Organization, authority sdk.ValAddress) MsgAppendOrganization {
	return MsgAppendOrganization{
		Organization: organization,
		Authority:    authority,
	}
}

const AppendOrganizationConst = "AppendOrganization"

// nolint
func (msg MsgAppendOrganization) Route() string { return RouterKey }
func (msg MsgAppendOrganization) Type() string  { return AppendOrganizationConst }
func (msg MsgAppendOrganization) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Authority)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAppendOrganization) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgAppendOrganization) ValidateBasic() error {
	if msg.Authority.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing authority address")
	}
	if msg.Organization.GetAddress().Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing organization address")
	}
	if msg.Organization.GetName() == "" {
		return sdkerrors.Wrap(ErrInvalidOrganization, "missing organization name")
	}
	return nil
}

/**
 * MsgChangeOrganizationApproval
 */
type MsgChangeOrganizationApproval struct {
	Organization sdk.AccAddress `json:"organization"`
	Authority    sdk.ValAddress `json:"authority"`
	Approve      bool           `json:approve"`
}

func NewMsgChangeOrganizationApproval(organization sdk.AccAddress, authority sdk.ValAddress, approve bool) MsgChangeOrganizationApproval {
	return MsgChangeOrganizationApproval{
		Organization: organization,
		Authority:    authority,
		Approve:      approve,
	}
}

const ChangeOrganizationApprovalConst = "ChangeOrganizationApproval"

// nolint
func (msg MsgChangeOrganizationApproval) Route() string { return RouterKey }
func (msg MsgChangeOrganizationApproval) Type() string  { return ChangeOrganizationApprovalConst }
func (msg MsgChangeOrganizationApproval) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Authority)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgChangeOrganizationApproval) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgChangeOrganizationApproval) ValidateBasic() error {
	if msg.Authority.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing authority address")
	}
	if msg.Organization.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing organization address")
	}
	return nil
}

/**
 * MsgCreateProduct
 */
type MsgCreateProduct struct {
	Product Product `json:"product"`
}

func NewMsgCreateProduct(product Product) MsgCreateProduct {
	return MsgCreateProduct{
		Product: product,
	}
}

const CreateProductConst = "CreateProduct"

// nolint
func (msg MsgCreateProduct) Route() string { return RouterKey }
func (msg MsgCreateProduct) Type() string  { return CreateProductConst }
func (msg MsgCreateProduct) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Product.GetManufacturer())}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateProduct) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateProduct) ValidateBasic() error {
	return msg.Product.Validate()
}

/**
 * MsgCreateUnit
 */
type MsgCreateUnit struct {
	Product    Product  `json:"product"`
	Details    string   `json:"details"`
	Components []string `json:"components"`
}

func NewMsgCreateUnit(product Product, details string, components []string) MsgCreateUnit {
	return MsgCreateUnit{
		Product:    product,
		Details:    details,
		Components: components,
	}
}

const CreateUnitConst = "CreateUnit"

// nolint
func (msg MsgCreateUnit) Route() string { return RouterKey }
func (msg MsgCreateUnit) Type() string  { return CreateUnitConst }
func (msg MsgCreateUnit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Product.GetManufacturer())}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateUnit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateUnit) ValidateBasic() error {
	err := msg.Product.Validate()
	if err != nil {
		return err
	}

	// Check all component references has the right length
	if len(msg.Components) > ComponentsMaxNumber {
		return sdkerrors.Wrap(ErrInvalidUnit, "too much components")
	}
	for _, componentReference := range msg.Components {
		if len(componentReference) != UnitReferenceLength {
			return sdkerrors.Wrap(ErrInvalidUnit, "a component reference is incorrect")
		}
	}

	return nil
}

/**
 * MsgTransferUnit
 */
type MsgTransferUnit struct {
	Unit      Unit           `json:"unit"`
	NewHolder sdk.AccAddress `json:"new_holder"`
}

func NewMsgTransferUnit(unit Unit, newHolder sdk.AccAddress) MsgTransferUnit {
	return MsgTransferUnit{
		Unit:      unit,
		NewHolder: newHolder,
	}
}

const TransferUnitConst = "TransferUnit"

// nolint
func (msg MsgTransferUnit) Route() string { return RouterKey }
func (msg MsgTransferUnit) Type() string  { return TransferUnitConst }
func (msg MsgTransferUnit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Unit.GetCurrentHolder())}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgTransferUnit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgTransferUnit) ValidateBasic() error {
	err := msg.Unit.Validate()
	if err != nil {
		return err
	}

	// Cannot transfer to himself
	if msg.Unit.GetCurrentHolder().Equals(msg.NewHolder) {
		return sdkerrors.Wrap(ErrInvalidTransfer, "cannot transfer to oneself")
	}

	return nil
}
