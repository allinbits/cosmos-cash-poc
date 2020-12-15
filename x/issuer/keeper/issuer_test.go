package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/allinbits/cosmos-cash-poc/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeperIssuerFunctions(t *testing.T) {
	ctx, keeper := MakeTestCtxAndKeeper(t)

	// Create test issuer
	valPubKey := MakeTestPubKey(SamplePubKey)
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	issuer := types.NewIssuer(
		"euro-token-issuer",
		accAddr,
		"euro",
	)

	// Set a value in the store
	keeper.SetIssuer(ctx, []byte(issuer.Name), issuer)

	// Check the store to see if the item was saved
	_, found := keeper.GetIssuer(ctx, []byte(issuer.Name))
	require.True(t, found)
}
