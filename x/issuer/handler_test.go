package issuer

import (
	"testing"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestHandleCreateIssuer(t *testing.T) {
	valPubKey := keeper.MakeTestPubKey(keeper.SamplePubKey)
	ctx, k := keeper.MakeTestCtxAndKeeper(t)

	name := "name"
	token := "nametoken"
	accAddr := sdk.AccAddress(valPubKey.Address().Bytes())

	msg := types.NewMsgCreateIssuer(
		name,
		accAddr,
		token,
		"10000",
		accAddr,
	)

	res, err := handleMsgCreateIssuer(ctx, msg, k)

	// Validate msg was handled correctly by checking issuer in store
	_, found := k.GetIssuer(ctx, accAddr)
	require.True(t, found)

	// No errors and res is populates
	require.NoError(t, err)
	require.NotNil(t, res)
}
