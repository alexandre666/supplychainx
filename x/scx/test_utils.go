package scx

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	poakeeper "github.com/ltacker/poa/keeper"
	poatypes "github.com/ltacker/poa/types"
	"github.com/ltacker/supplychainx/x/scx/keeper"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// This package contains various mocks for testing purpose

// Context and keeper used for mocking purpose
func MockContext() (sdk.Context, keeper.Keeper, poakeeper.Keeper) {
	// Store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, poatypes.StoreKey, params.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	cdc := codec.New()

	// Create the params keeper
	paramsKeeper := params.NewKeeper(cdc, keys[params.StoreKey], tKeys[params.TStoreKey])

	// Create a poa keeper
	poaKeeper := poakeeper.NewKeeper(cdc, keys[poatypes.StoreKey], paramsKeeper.Subspace(types.ModuleName))

	// Create a scx keeper
	scxKeeper := keeper.NewKeeper(cdc, keys[types.StoreKey], poaKeeper)

	// Create multiStore in memory
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	// Mount stores
	cms.MountStoreWithDB(keys[types.StoreKey], sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(keys[poatypes.StoreKey], sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(keys[params.StoreKey], sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(tKeys[params.TStoreKey], sdk.StoreTypeTransient, db)
	cms.LoadLatestVersion()

	// Create context
	ctx := sdk.NewContext(cms, abci.Header{}, false, log.NewNopLogger())

	return ctx, scxKeeper, poaKeeper
}

// Create a validator for test
func MockValidator() (poatypes.Validator, string) {
	// Junk description
	validatorDescription := poatypes.Description{
		Moniker:         "Moniker",
		Identity:        "Identity",
		Website:         "Website",
		SecurityContact: "SecurityContact",
		Details:         "Details",
	}

	// Generate operator address
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	operatorAddress := sdk.ValAddress(addr)

	// Generate a consPubKey
	pk = ed25519.GenPrivKey().PubKey()
	consPubKey, err := sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pk)
	if err != nil {
		panic(fmt.Sprintf("Cannot create a consPubKey: %v", err))
	}
	validator := poatypes.Validator{
		OperatorAddress: operatorAddress,
		ConsensusPubkey: consPubKey,
		Description:     validatorDescription,
	}

	return validator, consPubKey
}

// A misc organization
func MockOrganization() types.Organization {
	return types.NewOrganization(MockAccAddress(), "acme", "A company making everything")
}

// Create an validator address
func MockValAddress() sdk.ValAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.ValAddress(addr)
}

// Create an account address
func MockAccAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}

// Create a mock unit
func MockUnit(product types.Product, unitNumber uint, components []string) types.Unit {
	ref, _ := types.GetUnitReferenceFromProductAndUnitNumber(product.GetName(), unitNumber)
	unit := types.NewUnit(ref, product, "Perfectly manufactured", components)
	return unit
}
