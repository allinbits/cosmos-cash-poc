package regulator

import (
	"github.com/allinbits/cosmos-cash-poc/x/regulator/keeper"
	"github.com/allinbits/cosmos-cash-poc/x/regulator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data authtypes.GenesisState) {
	// To keep the module simple we set all genesis accounts as regualtors
	for _, account := range data.Accounts {
		reg := types.NewRegulator(account.GetAddress())
		k.SetRegulator(ctx, []byte(account.GetAddress()), reg)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	return types.NewGenesisState()
}
