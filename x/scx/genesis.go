package scx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ltacker/supplychainx/x/scx/keeper"
	"github.com/ltacker/supplychainx/x/scx/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	return types.NewGenesisState()
}
