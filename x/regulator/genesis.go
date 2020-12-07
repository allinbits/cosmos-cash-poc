package regulator

import (
	"fmt"
	"github.com/allinbits/cosmos-cash-poa/x/regulator/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/regulator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data authtypes.GenesisState) {
	// To keep the module simple we set all genesis accounts as regualtos
	fmt.Println(data)
	//k.SetRegulator()
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	return types.NewGenesisState()
}
