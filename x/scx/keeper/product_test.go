package keeper_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ltacker/supplychainx/x/scx"
	"github.com/ltacker/supplychainx/x/scx/types"
)

func TestAppendProduct(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "xphone", "A revolutionary phone")

	if product.GetUnitCount() != 0 {
		t.Errorf("unit count of a new product should be 0")
	}

	exist := scxKeeper.AppendProduct(ctx, product)
	if exist {
		t.Errorf("AppendProduct should not return true if the product doesn't exist")
	}
	exist = scxKeeper.AppendProduct(ctx, product)
	if !exist {
		t.Errorf("AppendProduct should return true if the product exists")
	}
}

func TestGetProduct(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "xphone", "A revolutionary phone")
	product2 := types.NewProduct(organization.GetAddress(), "xphone2", "A revolutionary phone")

	scxKeeper.AppendProduct(ctx, product)
	retrieved, found := scxKeeper.GetProduct(ctx, product.GetName())
	if !found {
		t.Errorf("GetProduct should find the product")
	}
	if !cmp.Equal(product, retrieved) {
		t.Errorf("GetProduct should find %v, found %v", product, retrieved)
	}

	// Should not find a unset product
	_, found = scxKeeper.GetProduct(ctx, product2.GetName())
	if found {
		t.Errorf("GetProduct should not find not appended product")
	}
}

func TestIncreaseProductCount(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "xphone", "A revolutionary phone")

	// Not existing product
	found := scxKeeper.IncreaseProductCount(ctx, product.GetName())
	if found {
		t.Errorf("IncreaseProductCount should not find not appended product")
	}

	scxKeeper.AppendProduct(ctx, product)

	// Increment the count
	found = scxKeeper.IncreaseProductCount(ctx, product.GetName())
	if !found {
		t.Errorf("IncreaseProductCount should find appended product")
	}
	retrieved, _ := scxKeeper.GetProduct(ctx, product.GetName())
	if retrieved.GetUnitCount() != 1 {
		t.Errorf("IncreaseProductCount should increase product unit count")
	}
}

func TestGetAllProducts(t *testing.T) {
	ctx, scxKeeper, _ := scx.MockContext()
	organization := scx.MockOrganization()
	product := types.NewProduct(organization.GetAddress(), "xphone", "A revolutionary phone")
	product2 := types.NewProduct(organization.GetAddress(), "xphone2", "A revolutionary phone")

	scxKeeper.AppendProduct(ctx, product)
	scxKeeper.AppendProduct(ctx, product2)

	retrievedProducts := scxKeeper.GetAllProducts(ctx)
	if len(retrievedProducts) != 2 {
		t.Errorf("GetAllProducts should find %v validators, found %v", 2, len(retrievedProducts))
	}
}
