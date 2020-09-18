package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/issuer/issuers",
		getAllIssuersHandlerFn(cliCtx),
	).Methods("GET")
}

func getAllIssuersHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cdc := codec.New()
		codec.RegisterCrypto(cdc)
		cliCtx = cliCtx.WithCodec(cdc)

		resKVs, height, err := cliCtx.QuerySubspace(types.IssuersKey, types.RouterKey)
		if err != nil {
			panic(err)
		}

		var issuers []types.Issuer
		for _, kv := range resKVs {
			issuer := types.Issuer{}
			cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &issuer)
			issuers = append(issuers, issuer)

		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, issuers)
	}
}
