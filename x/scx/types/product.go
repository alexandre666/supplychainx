package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Describes a product
type Product struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Manufacturer sdk.AccAddress `json:"manufacturer"`
	UnitCount    uint64         `json:"count"`
}

func NewProduct(m sdk.AccAddress, name, description string) Product {
	return Product{
		Manufacturer: m,
		Name:         name,
		Description:  description,
		UnitCount:    0,
	}
}

func (p Product) GetManufacturer() sdk.AccAddress {
	return p.Manufacturer
}

func (p Product) GetName() string {
	return p.Name
}

func (p Product) GetDescription() string {
	return p.Description
}

func (p Product) GetUnitCount() uint64 {
	return p.UnitCount
}

func (p *Product) IncreaseUnit() {
	p.UnitCount += 1
}

func (p Product) Validate() error {
	if p.GetManufacturer().Empty() {
		return sdkerrors.Wrap(ErrInvalidProduct, "missing manufacturer")
	}
	if p.GetName() == "" {
		return sdkerrors.Wrap(ErrInvalidProduct, "missing product name")
	}

	return nil
}

// Encoding functions
func MustMarshalProduct(cdc *codec.Codec, p Product) []byte {
	return cdc.MustMarshalBinaryBare(&p)
}
func MustUnmarshalProduct(cdc *codec.Codec, value []byte) Product {
	p, err := UnmarshalProduct(cdc, value)
	if err != nil {
		panic(p)
	}

	return p
}
func UnmarshalProduct(cdc *codec.Codec, value []byte) (p Product, err error) {
	err = cdc.UnmarshalBinaryBare(value, &p)
	return p, err
}
