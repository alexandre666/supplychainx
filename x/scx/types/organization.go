package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Organizations are the entities that can create new products and units
type Organization struct {
	Address     sdk.AccAddress `json:"address"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Approved    bool           `json:"approved"`
}

func NewOrganization(address sdk.AccAddress, name, description string) Organization {
	return Organization{
		Address:     address,
		Name:        name,
		Description: description,
		Approved:    true,
	}
}

func (o Organization) GetAddress() sdk.AccAddress {
	return o.Address
}

func (o Organization) GetName() string {
	return o.Name
}

func (o Organization) GetDescription() string {
	return o.Description
}

func (o Organization) IsApproved() bool {
	return o.Approved
}

func (o *Organization) Approve() {
	o.Approved = true
}

func (o *Organization) Relegate() {
	o.Approved = false
}

// Encoding functions
func MustMarshalOrganization(cdc *codec.Codec, o Organization) []byte {
	return cdc.MustMarshalBinaryBare(&o)
}
func MustUnmarshalOrganization(cdc *codec.Codec, value []byte) Organization {
	org, err := UnmarshalOrganization(cdc, value)
	if err != nil {
		panic(err)
	}

	return org
}
func UnmarshalOrganization(cdc *codec.Codec, value []byte) (o Organization, err error) {
	err = cdc.UnmarshalBinaryBare(value, &o)
	return o, err
}
