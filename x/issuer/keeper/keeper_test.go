package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeperGenericFunctions(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// Create test entity
	valPubKey := MakeTestPubKey(SamplePubKey)
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	entity := types.NewIssuer(
		"euro-token-issuer",
		accAddr,
		"euro",
	)

	// Set a value in the store
	keeper.Set(ctx, []byte(entity.Name), types.IssuersKey, entity)

	// Check the store to see if the item was saved
	_, found := keeper.Get(ctx, []byte(entity.Name), types.IssuersKey, keeper.UnmarshalIssuer)
	require.True(t, found)

	// Get all items from the store
	allVals := keeper.GetAll(ctx, types.IssuersKey, keeper.UnmarshalIssuer)
	require.Equal(t, 1, len(allVals))
}
