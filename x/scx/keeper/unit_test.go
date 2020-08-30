package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/google/go-cmp/cmp"
	"github.com/ltacker/supplychainx/x/scx"
	"github.com/ltacker/supplychainx/x/scx/types"
)

func TestAppendUnit(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "x", "")
	unit := scx.MockUnit(product, 0, []string{})
	unit2 := scx.MockUnit(product, 1, []string{})

	// Can append an unit
	exist := scxKeeper.AppendUnit(ctx, unit)
	if exist {
		t.Errorf("AppendUnit should not return true if the unit doesn't exist")
	}
	exist = scxKeeper.AppendUnit(ctx, unit)
	if !exist {
		t.Errorf("AppendUnit should return true if the unit exists")
	}
	exist = scxKeeper.AppendUnit(ctx, unit2)
	if exist {
		t.Errorf("AppendUnit can append other units")
	}
}

func TestGetUnit(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "x", "")
	unit := scx.MockUnit(product, 0, []string{})
	unit2 := scx.MockUnit(product, 1, []string{})

	scxKeeper.AppendUnit(ctx, unit)
	retrieved, found := scxKeeper.GetUnit(ctx, unit.GetReference())
	if !found {
		t.Errorf("GetUnit should find the unit")
	}
	if !cmp.Equal(unit.GetReference(), retrieved.GetReference()) {
		t.Errorf("GetUnit should find %v, found %v", unit, retrieved)
	}

	// Should not find a unset unit
	_, found = scxKeeper.GetUnit(ctx, unit2.GetReference())
	if found {
		t.Errorf("GetUnit should not find not appended unit")
	}
}

func TestGetUnitTrace(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization1 := scx.MockOrganization()
	organization2 := scx.MockOrganization()
	organization3 := scx.MockOrganization()
	organization4 := scx.MockOrganization()
	product := types.NewProduct(organization1.GetAddress(), "x", "")
	unit := scx.MockUnit(product, 0, []string{})

	// Transfer unit
	unit.ChangeHolder(organization2.GetAddress())
	unit.ChangeHolder(organization3.GetAddress())
	unit.ChangeHolder(organization4.GetAddress())
	scxKeeper.SetUnit(ctx, unit)

	// Check the tract
	trace, _ := scxKeeper.GetUnitTrace(ctx, unit.GetReference())
	expectedTrace := []sdk.AccAddress{organization1.GetAddress(), organization2.GetAddress(), organization3.GetAddress(), organization4.GetAddress()}
	if !cmp.Equal(trace, expectedTrace) {
		t.Errorf("GetUnitTrace should find %v, found %v", expectedTrace, trace)
	}
}

func TestGetUnitComponents(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization1 := scx.MockOrganization()
	product := types.NewProduct(organization1.GetAddress(), "x", "")
	unit1 := scx.MockUnit(product, 0, []string{})
	unit2 := scx.MockUnit(product, 1, []string{})
	unit3 := scx.MockUnit(product, 2, []string{unit1.GetReference(), unit2.GetReference()})
	unit4 := scx.MockUnit(product, 3, []string{})
	unit5 := scx.MockUnit(product, 4, []string{unit3.GetReference()})
	unit6 := scx.MockUnit(product, 5, []string{unit4.GetReference(), unit5.GetReference()})

	scxKeeper.SetUnit(ctx, unit1)
	scxKeeper.SetUnit(ctx, unit2)
	scxKeeper.SetUnit(ctx, unit3)
	scxKeeper.SetUnit(ctx, unit4)
	scxKeeper.SetUnit(ctx, unit5)
	scxKeeper.SetUnit(ctx, unit6)

	// Get all the components
	components, _ := scxKeeper.GetUnitComponents(ctx, unit6.GetReference())

	if len(components) != 5 {
		t.Errorf("GetUnitComponents should find 5 components")
	}

	componentRefs := []string{components[0].GetReference(), components[1].GetReference(), components[2].GetReference(), components[3].GetReference(), components[4].GetReference()}
	expectedComponentRefs := []string{unit4.GetReference(), unit5.GetReference(), unit3.GetReference(), unit1.GetReference(), unit2.GetReference()}

	if !cmp.Equal(componentRefs, expectedComponentRefs) {
		t.Errorf("GetUnitComponents should find %v, found %v", expectedComponentRefs, componentRefs)
	}
}
