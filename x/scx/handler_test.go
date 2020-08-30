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
	// TODO

	// Cannot create if the manufacturer doesn't exist

	// Cannot create if the manufacturer is relegated

	// Cannot create if the product doesn't exist

	// Cannot create if a component doesn't exist

	// Cannot create if a component is not owned

	// Cannot create if a component is already component of another unit

	// Create a new unit

	// The components' "componentOf" field is updated

	// The product count is incremented
}

func TestHandleMsgTransferUnit(t *testing.T) {
	// TODO

	// Cannot transfer if the holder doesn't exist

	// Cannot transfer if the holder is relegated

	// Cannot transfer if the new holder doesn't exist

	// Cannot transfer if the new holder is relegated

	// Cannot transfer if the unit doesn't exist

	// Cannot transfer if the holder doesn't own the unit

	// Cannot transfer if the unit is a component of another unit

	// Can transfer the unit

	// The holder is update

	// The holder history is updated
}
