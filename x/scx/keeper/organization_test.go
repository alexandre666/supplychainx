package keeper_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ltacker/supplychainx/x/scx"
)

func TestGetOrganization(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	organizationNotSet := scx.MockOrganization()

	scxKeeper.SetOrganization(ctx, organization)

	// Should find the correct organization
	retrievedOrganization, found := scxKeeper.GetOrganization(ctx, organization.GetAddress())
	if !found {
		t.Errorf("GetOrganization should find organization if it has been set")
	}

	if !cmp.Equal(organization, retrievedOrganization) {
		t.Errorf("GetValidator should find %v, found %v", organization, retrievedOrganization)
	}

	// Should not find a unset organization
	_, found = scxKeeper.GetOrganization(ctx, organizationNotSet.GetAddress())
	if found {
		t.Errorf("GetOrganization should not find organization if it has not been set")
	}
}

func TestGetAllOrganizations(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization1 := scx.MockOrganization()
	organization2 := scx.MockOrganization()

	scxKeeper.SetOrganization(ctx, organization1)
	scxKeeper.SetOrganization(ctx, organization2)

	retrievedOrganizations := scxKeeper.GetAllOrganizations(ctx)
	if len(retrievedOrganizations) != 2 {
		t.Errorf("GetAllOrganizations should find %v validators, found %v", 2, len(retrievedOrganizations))
	}
}
