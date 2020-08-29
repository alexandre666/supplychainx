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
}

func NewOrganization(address sdk.AccAddress, name, description string) Organization {
	return Organization{
		Address:     address,
		Name:        name,
		Description: description,
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

// Encoding functions
func MustMarshalVote(cdc *codec.Codec, o Organization) []byte {
	return cdc.MustMarshalBinaryBare(&o)
}
func MustUnmarshalVote(cdc *codec.Codec, value []byte) Organization {
	org, err := UnmarshalVote(cdc, value)
	if err != nil {
		panic(err)
	}

	return org
}
func UnmarshalVote(cdc *codec.Codec, value []byte) (o Organization, err error) {
	err = cdc.UnmarshalBinaryBare(value, &o)
	return o, err
}
