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
		t.Errorf("MsgAppendOrganization should append an organisation, got error %v", err)
	}
	_, found := scxKeeper.GetOrganization(ctx, organization.GetAddress())
	if !found {
		t.Errorf("MsgAppendOrganization should append an organisation, not found in keeper")
	}

	// Cannot append an existing organization
	organization.Name = "Acme2"
	msg = types.NewMsgAppendOrganization(organization, authority)
	_, err = handler(ctx, msg)
	if err.Error() != types.ErrOrganizationAlreadyExists.Error() {
		t.Errorf("MsgAppendOrganization with existing organization, error should be %v, got %v", types.ErrOrganizationAlreadyExists.Error(), err.Error())
	}
}
