package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allinbits/cosmos-cash-poa/x/did/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// TODO: Define your GET REST endpoints
	r.HandleFunc(
		"/did/documents",
		getAllDidDocumentsHandlerFn(cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/did/creds",
		getAllCredsHandlerFn(cliCtx),
	).Methods("GET")
}

func getAllDidDocumentsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cdc := codec.New()
		codec.RegisterCrypto(cdc)
		cliCtx = cliCtx.WithCodec(cdc)

		resKVs, height, err := cliCtx.QuerySubspace(types.DidDocumentKey, "did")
		if err != nil {
			panic(err)
		}

		var diddocuments []types.DidDocument
		for _, kv := range resKVs {
			doc := types.DidDocument{}
			cdc.UnmarshalBinaryBare(kv.Value, &doc)
			diddocuments = append(diddocuments, doc)

		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, diddocuments)
	}
}

func getAllCredsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cdc := codec.New()
		codec.RegisterCrypto(cdc)
		cliCtx = cliCtx.WithCodec(cdc)

		resKVs, height, err := cliCtx.QuerySubspace(types.VerifiableCredentialKey, "did")
		if err != nil {
			panic(err)
		}

		var verifiablecreds []types.VerifiableCredential
		for _, kv := range resKVs {
			vc := types.VerifiableCredential{}
			cdc.UnmarshalBinaryBare(kv.Value, &vc)
			verifiablecreds = append(verifiablecreds, vc)

		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, verifiablecreds)
	}
}
