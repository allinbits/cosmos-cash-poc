package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/allinbits/cosmos-cash-poa/x/poa/types"
)

func TestKeeperParamsFunctions(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// SetParams test
	keeper.SetParams(ctx, types.DefaultParams())

	// GetParams test
	params := keeper.GetParams(ctx)

	require.Equal(t, uint16(50), params.Quorum)
	require.Equal(t, uint16(100), params.MaxValidators)
}
