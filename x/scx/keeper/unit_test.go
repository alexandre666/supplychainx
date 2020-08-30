package keeper

import (
	"testing"
)

// Append an unit
func TestAppendUnit(t *testing.T) {
	// TODO

	// Can append an unit

	// Append an existing unit return false
}

// Fetch a unit from its reference
func TestGetUnit(t *testing.T) {
	// TODO

	// Get an unit

	// Non existing unit return false
}

// Get the list of all the holder
func TestGetUnitTrace(t *testing.T) {
	// TODO

	// Get all the addresses
}

// Get all the component units of the unit
func TestGetUnitComponents(t *testing.T) {
	// TODO

	// Get all the components
}

// Transfer the ownership of an unit
func TestTransferUnit(t *testing.T) {
	// TODO

	// Cannot transfer if the unit is already component of another unit
}

// Set the ComponentOf field
func TestSetComponentOf(t *testing.T) {
	// TODO

	// Cannot set if if the unit is already component of another unit
}
