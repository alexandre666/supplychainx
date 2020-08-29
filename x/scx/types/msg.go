package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var _ sdk.Msg = &MsgAppendOrganization{}

type MsgAppendOrganization struct {
	Organization Organization   `json:"organization"`
	Authority    sdk.ValAddress `json:"aauthority"`
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
