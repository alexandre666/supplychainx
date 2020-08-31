package scx_test

import (
	"testing"

	"github.com/ltacker/supplychainx/x/scx"
	"github.com/ltacker/supplychainx/x/scx/types"
)

func TestHandleMsgAppendOrganization(t *testing.T) {
	ctx, scxKeeper, poaKeeper := scx.MockContext()
	organization := scx.MockOrganization()
	validator, _ := scx.MockValidator()
	authority := validator.GetOperator()
	notAnAuthority := scx.MockValAddress()

	handler := scx.NewHandler(scxKeeper)

	// Create validator
	poaKeeper.AppendValidator(ctx, validator)

	// A non authority cannot add a new organization
	msg := types.NewMsgAppendOrganization(organization, notAnAuthority)
	_, err := handler(ctx, msg)
	if err.Error() != types.ErrNotAnAuthority.Error() {
		t.Errorf("MsgAppendOrganization sender not an authority, error should be %v, got %v", types.ErrNotAnAuthority.Error(), err.Error())
	}

	// An authority can add a new organization
	msg = types.NewMsgAppendOrganization(organization, authority)
	_, err = handler(ctx, msg)
	if err != nil {
		t.Errorf("MsgAppendOrganization should append an organization, got error %v", err)
	}
	organization, found := scxKeeper.GetOrganization(ctx, organization.GetAddress())
	if !found {
		t.Errorf("MsgAppendOrganization should append an organization, not found in keeper")
	}
	if !organization.IsApproved() {
		t.Errorf("MsgAppendOrganization a new organization should be approved")
	}

	// Cannot append an existing organization
	organization.Name = "Acme2"
	msg = types.NewMsgAppendOrganization(organization, authority)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationAlreadyExists.Error() {
		t.Errorf("MsgAppendOrganization with existing organization, error should be %v, got %v", types.ErrOrganizationAlreadyExists.Error(), err.Error())
	}
}

func TestHandleMsgChangeOrganizationApproval(t *testing.T) {
	ctx, scxKeeper, poaKeeper := scx.MockContext()
	organization := scx.MockOrganization()
	validator, _ := scx.MockValidator()
	authority := validator.GetOperator()
	notAnAuthority := scx.MockValAddress()

	handler := scx.NewHandler(scxKeeper)

	// Create validator
	poaKeeper.AppendValidator(ctx, validator)

	// Cannot update a non existing organization
	msg := types.NewMsgChangeOrganizationApproval(organization.GetAddress(), authority, false)
	_, err := handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotFound.Error() {
		t.Errorf("MsgChangeOrganizationApproval with non existing organization, error should be %v, got %v", types.ErrOrganizationNotFound.Error(), err.Error())
	}

	// Create organization
	scxKeeper.SetOrganization(ctx, organization)

	// A non authority cannot update an organization
	msg = types.NewMsgChangeOrganizationApproval(organization.GetAddress(), notAnAuthority, false)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrNotAnAuthority.Error() {
		t.Errorf("MsgChangeOrganizationApproval sender not an authority, error should be %v, got %v", types.ErrNotAnAuthority.Error(), err.Error())
	}

	// Cannot reapprove an already approved organization
	msg = types.NewMsgChangeOrganizationApproval(organization.GetAddress(), authority, true)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationAlreadyApproved.Error() {
		t.Errorf("ReapproveOrganization organization already approved, error should be %v, got %v", types.ErrOrganizationAlreadyApproved.Error(), err.Error())
	}

	// Relegate an approved organization
	msg = types.NewMsgChangeOrganizationApproval(organization.GetAddress(), authority, false)
	_, err = handler(ctx, msg)
	if err != nil {
		t.Errorf("RelegateOrganization unexpected error %v", err.Error())
	}
	organizationFound, _ := scxKeeper.GetOrganization(ctx, organization.GetAddress())
	if organizationFound.IsApproved() {
		t.Errorf("RelegateOrganization should relegate organization")
	}

	// Cannot relegate an already relegated organization
	msg = types.NewMsgChangeOrganizationApproval(organization.GetAddress(), authority, false)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationAlreadyRelegated.Error() {
		t.Errorf("RelegateOrganization organization already relegated, error should be %v, got %v", types.ErrOrganizationAlreadyRelegated.Error(), err.Error())
	}

	// Reapprove an organization
	msg = types.NewMsgChangeOrganizationApproval(organization.GetAddress(), authority, true)
	_, err = handler(ctx, msg)
	if err != nil {
		t.Errorf("ReapproveOrganization unexpected error %v", err.Error())
	}
	organizationFound, _ = scxKeeper.GetOrganization(ctx, organization.GetAddress())
	if !organizationFound.IsApproved() {
		t.Errorf("ReapproveOrganization should reapprove organization")
	}
}

func TestHandleMsgCreateProduct(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	notAnOrganization := scx.MockAccAddress()
	organization := scx.MockOrganization()

	handler := scx.NewHandler(scxKeeper)

	// Create Organization
	scxKeeper.SetOrganization(ctx, organization)

	// A non organization cannot create a product
	product := types.NewProduct(notAnOrganization, "xphone", "A revolutionary phone")
	msg := types.NewMsgCreateProduct(product)
	_, err := handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotFound.Error() {
		t.Errorf("MsgCreateProduct not existing organization, error should be %v, got %v", types.ErrOrganizationNotFound.Error(), err.Error())
	}

	// An approved organization can create a product
	product = types.NewProduct(organization.GetAddress(), "xphone", "A revolutionary phone")
	msg = types.NewMsgCreateProduct(product)
	_, err = handler(ctx, msg)
	if err != nil {
		t.Errorf("MsgCreateProduct existing organization, unexpected error %v", err.Error())
	}
	_, found := scxKeeper.GetProduct(ctx, product.GetName())
	if !found {
		t.Errorf("MsgCreateProduct, created product not found")
	}

	// Cannot create an already existing product
	product = types.NewProduct(organization.GetAddress(), "xphone", "A copy of the revolutionary phone")
	msg = types.NewMsgCreateProduct(product)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrProductAlreadyExist.Error() {
		t.Errorf("MsgCreateProduct existing product, error should be %v, got %v", types.ErrProductAlreadyExist.Error(), err.Error())
	}

	// A relegated organization cannot create new products
	organization.Relegate()
	scxKeeper.SetOrganization(ctx, organization)
	product2 := types.NewProduct(organization.GetAddress(), "xphone2", "A new revolutionary phone")
	msg = types.NewMsgCreateProduct(product2)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotApproved.Error() {
		t.Errorf("MsgCreateProduct from relegated organization, error should be %v, got %v", types.ErrOrganizationNotApproved.Error(), err.Error())
	}
}

func TestHandleMsgCreateUnit(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	orgAddress := organization.GetAddress()
	product := types.NewProduct(orgAddress, "xphone", "A revolutionary phone")
	productName := product.GetName()
	handler := scx.NewHandler(scxKeeper)

	scxKeeper.AppendProduct(ctx, product)

	// Cannot create if the manufacturer doesn't exist
	notAnOrganization := scx.MockAccAddress()
	msg := types.NewMsgCreateUnit(productName, notAnOrganization, "", []string{})
	_, err := handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotFound.Error() {
		t.Errorf("MsgCreateUnit with non existing manufacturer should fail: %v", err)
	}

	// Cannot create if the manufacturer is relegated
	organization.Relegate()
	scxKeeper.SetOrganization(ctx, organization)
	msg = types.NewMsgCreateUnit(productName, orgAddress, "", []string{})
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotApproved.Error() {
		t.Errorf("MsgCreateUnit with relegated manufacturer should fail with %v, got %v", types.ErrOrganizationNotApproved, err)
	}

	organization.Approve()
	scxKeeper.SetOrganization(ctx, organization)

	// Cannot create if the manufacturer is not the manufacturer of the product
	otherProduct := types.NewProduct(notAnOrganization, "toast", "")
	scxKeeper.AppendProduct(ctx, otherProduct)
	msg = types.NewMsgCreateUnit(otherProduct.GetName(), orgAddress, "", []string{})
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrInvalidOrganization.Error() {
		t.Errorf("MsgCreateUnit with invalid manufacturer should fail: %v", err)
	}

	// Cannot create if the product doesn't exist
	inexistantProduct := types.NewProduct(orgAddress, "nothing", "")
	msg = types.NewMsgCreateUnit(inexistantProduct.GetName(), orgAddress, "", []string{})
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrProductNotFound.Error() {
		t.Errorf("MsgCreateUnit with non existing product should fail: %v", err)
	}

	// Cannot create if a component doesn't exist
	inexistantRef := "aaaaaaaaaaaaaaaa"
	msg = types.NewMsgCreateUnit(productName, orgAddress, "", []string{inexistantRef})
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrUnitNotFound.Error() {
		t.Errorf("MsgCreateUnit with non existing component should fail: %v", err)
	}

	// Cannot create if a component is not owned
	notOwnedUnit := scx.MockUnit(otherProduct, 100, []string{})
	scxKeeper.SetUnit(ctx, notOwnedUnit)
	msg = types.NewMsgCreateUnit(productName, orgAddress, "", []string{notOwnedUnit.GetReference()})
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrComponentNotOwned.Error() {
		t.Errorf("MsgCreateUnit with not owned component should fail: %v", err)
	}

	// Cannot create if a component is already component of another unit
	unitAlreadyComponent := scx.MockUnit(product, 100, []string{})
	unitAlreadyComponent.ComponentOf = "aaaaaaaaaaaaaaaa"
	scxKeeper.SetUnit(ctx, unitAlreadyComponent)
	msg = types.NewMsgCreateUnit(productName, orgAddress, "", []string{unitAlreadyComponent.GetReference()})
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrUnitComponentOfAnotherUnit.Error() {
		t.Errorf("MsgCreateUnit with a component already used should fail: %v", err)
	}

	// Can create a new unit
	component1 := scx.MockUnit(product, 101, []string{})
	component2 := scx.MockUnit(product, 102, []string{})
	component3 := scx.MockUnit(product, 103, []string{})
	scxKeeper.SetUnit(ctx, component1)
	scxKeeper.SetUnit(ctx, component2)
	scxKeeper.SetUnit(ctx, component3)
	msg = types.NewMsgCreateUnit(productName, orgAddress, "", []string{
		component1.GetReference(),
		component2.GetReference(),
		component3.GetReference(),
	})
	result, err := handler(ctx, msg)
	if err != nil {
		t.Errorf("MsgCreateUnit unexpected error creating a unit: %v", err)
	}
	var reference string
	types.ModuleCdc.MustUnmarshalBinaryLengthPrefixed(result.Data, &reference)

	_, found := scxKeeper.GetUnit(ctx, reference)
	if !found {
		t.Errorf("MsgCreateUnit didn't create a unit of reference: %v", reference)
	}

	// The components' "componentOf" field is updated
	component1, _ = scxKeeper.GetUnit(ctx, component1.GetReference())
	component2, _ = scxKeeper.GetUnit(ctx, component2.GetReference())
	component3, _ = scxKeeper.GetUnit(ctx, component3.GetReference())
	if component1.GetComponentOf() != reference || component2.GetComponentOf() != reference || component3.GetComponentOf() != reference {
		t.Errorf("MsgCreateUnit didn't update ComponentOf field of the components")
	}

	// The product count is incremented
	product, _ = scxKeeper.GetProduct(ctx, productName)
	if product.GetUnitCount() != 1 {
		t.Errorf("MsgCreateUnit should increase unit number of the product")
	}
}

func TestHandleMsgTransferUnit(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	orgAddress := organization.GetAddress()
	organization2 := scx.MockOrganization()
	orgAddress2 := organization2.GetAddress()
	product := types.NewProduct(orgAddress, "xphone", "A revolutionary phone")
	handler := scx.NewHandler(scxKeeper)
	unit := scx.MockUnit(product, 0, []string{})
	ref := unit.GetReference()

	scxKeeper.AppendProduct(ctx, product)
	scxKeeper.AppendUnit(ctx, unit)

	// Cannot transfer if the holder doesn't exist
	notAnOrganization := scx.MockAccAddress()
	msg := types.NewMsgTransferUnit(ref, notAnOrganization, orgAddress2)
	_, err := handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotFound.Error() {
		t.Errorf("MsgTransferUnit with non existing holder should fail: %v", err)
	}

	// Cannot transfer if the holder is relegated
	organization.Relegate()
	scxKeeper.SetOrganization(ctx, organization)
	msg = types.NewMsgTransferUnit(ref, orgAddress, orgAddress2)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotApproved.Error() {
		t.Errorf("MsgTransferUnit with relegated holder should fail: %v", err)
	}

	// Cannot transfer if the new holder doesn't exist
	organization.Approve()
	scxKeeper.SetOrganization(ctx, organization)
	msg = types.NewMsgTransferUnit(ref, orgAddress, notAnOrganization)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotFound.Error() {
		t.Errorf("MsgTransferUnit with non existing new holder should fail: %v", err)
	}

	// Cannot transfer if the new holder is relegated
	organization2.Relegate()
	scxKeeper.SetOrganization(ctx, organization2)
	msg = types.NewMsgTransferUnit(ref, orgAddress, orgAddress2)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationNotApproved.Error() {
		t.Errorf("MsgTransferUnit with relegated new holder should fail: %v", err)
	}

	// Cannot transfer if the unit doesn't exist
	organization2.Approve()
	scxKeeper.SetOrganization(ctx, organization2)
	notExistingUnit := scx.MockUnit(product, 1, []string{})
	msg = types.NewMsgTransferUnit(notExistingUnit.GetReference(), orgAddress, orgAddress2)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrUnitNotFound.Error() {
		t.Errorf("MsgTransferUnit with non existing unit should fail : %v", err)
	}

	// Cannot transfer if the holder doesn't own the unit
	unit.ChangeHolder(notAnOrganization)
	scxKeeper.SetUnit(ctx, unit)
	msg = types.NewMsgTransferUnit(ref, orgAddress, orgAddress2)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrUnitNotOwned.Error() {
		t.Errorf("MsgTransferUnit with non owned unit should fail: %v", err)
	}

	// Cannot transfer if the unit is a component of another unit
	unit.ChangeHolder(orgAddress)
	unit.ComponentOf = "aaaaaaaaaaaaaaaa"
	scxKeeper.SetUnit(ctx, unit)
	msg = types.NewMsgTransferUnit(ref, orgAddress, orgAddress2)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrUnitComponentOfAnotherUnit.Error() {
		t.Errorf("MsgTransferUnit with unit component of another unit should fail: %v", err)
	}

	// Can transfer the unit
	unit.ComponentOf = ""
	scxKeeper.SetUnit(ctx, unit)
	msg = types.NewMsgTransferUnit(ref, orgAddress, orgAddress2)
	_, err = handler(ctx, msg)
	if err != nil {
		t.Errorf("MsgTransferUnit unexpected error: %v", err)
	}

	// The holder is update
	retrieved, _ := scxKeeper.GetUnit(ctx, ref)
	if !retrieved.GetCurrentHolder().Equals(orgAddress2) {
		t.Errorf("MsgTransferUnit new holder should be %v, got %v", orgAddress2, retrieved.GetCurrentHolder())
	}
}
