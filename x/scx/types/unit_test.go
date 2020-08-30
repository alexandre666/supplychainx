package types_test

import (
	"testing"

	"github.com/ltacker/supplychainx/x/scx"
	"github.com/ltacker/supplychainx/x/scx/types"
)

func TestChangeHolder(t *testing.T) {
	organization := scx.MockOrganization()
	organization2 := scx.MockOrganization()
	organization3 := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "xphone", "A revolutionary phone")
	unit := scx.MockUnit(product, 0, []string{})

	if !unit.GetCurrentHolder().Equals(organization.GetAddress()) {
		t.Errorf("ChangeHolder, the manufacturer is not the original holder")
	}

	// Change holder
	unit.ChangeHolder(organization2.GetAddress())
	if !unit.GetCurrentHolder().Equals(organization2.GetAddress()) {
		t.Errorf("ChangeHolder, the holder has not been changed to organization2")
	}

	// Change holder once again
	unit.ChangeHolder(organization3.GetAddress())
	if !unit.GetCurrentHolder().Equals(organization3.GetAddress()) {
		t.Errorf("ChangeHolder, the holder has not been changed to organization3")
	}

	// Check holder history
	holderHistory := unit.GetHolderHistrory()
	if (len(holderHistory) != 2) || (!holderHistory[0].Equals(organization.GetAddress())) || (!holderHistory[1].Equals(organization2.GetAddress())) {
		t.Errorf("ChangeHolder, holder history is incorrect")
	}
}
