package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// The size of a unit reference (here the 16 first bytes of the unit hash)
	UnitReferenceLength = 32

	// Maximum number of components in a unit
	ComponentsMaxNumber = 50
)

// Describes the unit of a product
type Unit struct {
	Reference     string           `json:"reference"`
	Product       string           `json:"product"`
	Details       string           `json:"details"`
	Components    []string         `json:"components"`
	Holder        sdk.AccAddress   `json:"holder"`
	HolderHistory []sdk.AccAddress `json:"holder_history"`
	ComponentOf   string           `json:"component_of"`
}

func NewUnit(reference string, product Product, details string, components []string) Unit {
	return Unit{
		Reference:     reference,
		Product:       product.GetName(),
		Details:       details,
		Components:    components,
		Holder:        product.GetManufacturer(),
		HolderHistory: []sdk.AccAddress{},
		ComponentOf:   "",
	}
}

// TODO: Other functions

func (u Unit) GetCurrentHolder() sdk.AccAddress {
	return u.Holder
}

func (u Unit) Validate() error {
	if len(u.Holder) != UnitReferenceLength {
		return sdkerrors.Wrap(ErrInvalidUnit, "invalid reference")
	}
	if u.Holder.Empty() {
		return sdkerrors.Wrap(ErrInvalidUnit, "missing holder")
	}
	if u.Product == "" {
		return sdkerrors.Wrap(ErrInvalidUnit, "missing product name")
	}

	// Check all component references has the right length
	if len(u.Components) > ComponentsMaxNumber {
		return sdkerrors.Wrap(ErrInvalidUnit, "too much components")
	}
	for _, componentReference := range u.Components {
		if len(componentReference) != UnitReferenceLength {
			return sdkerrors.Wrap(ErrInvalidUnit, "a component reference is incorrect")
		}
	}

	// Check component of
	if (u.ComponentOf != "") && (len(u.ComponentOf) != UnitReferenceLength) {
		return sdkerrors.Wrap(ErrInvalidUnit, "component of is incorrect")
	}

	return nil
}

// Get the unit reference from the product name and its number
// The reference are the first bytes of the sha 256 hash of the product name associated with the unit number
func GetUnitReferenceFromProductAndUnitNumber(productName string, unitNumber uint) (string, error) {
	chunk := struct {
		ProductName string
		UnitNumber  uint
	}{productName, unitNumber}

	encodedChunk, err := json.Marshal(chunk)
	if err != nil {
		return "", err
	}

	// Compute the hash
	hash := sha256.Sum256(encodedChunk)

	// Get the reference
	byteLength := UnitReferenceLength / 2 // One byte = two chars
	referenceBytes := hash[:byteLength]
	return hex.EncodeToString(referenceBytes), nil
}
