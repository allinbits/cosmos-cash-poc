package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/allinbits/cosmos-cash-poa/x/did/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group did queries under a subcommand
	didQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	didQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdDidDocumentAll(queryRoute, cdc),
			GetCmdVerifiableCredentialAll(queryRoute, cdc),
		)...,
	)

	return didQueryCmd
}

func GetCmdDidDocumentAll(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "did-documents",
		Short: "did-documents",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.DidDocumentKey, "did")
			if err != nil {
				return err
			}

			var diddocuments []types.DidDocument
			for _, kv := range resKVs {
				doc := types.DidDocument{}
				cdc.UnmarshalBinaryBare(kv.Value, &doc)
				diddocuments = append(diddocuments, doc)

			}

			return cliCtx.PrintOutput(diddocuments)
		},
	}
}

func GetCmdVerifiableCredentialAll(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "verifiable-credentials",
		Short: "verifiable-credentials",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.VerifiableCredentialKey, "did")
			if err != nil {
				return err
			}

			var verifiablecreds []types.VerifiableCredential
			for _, kv := range resKVs {
				vc := types.VerifiableCredential{}
				cdc.UnmarshalBinaryBare(kv.Value, &vc)
				verifiablecreds = append(verifiablecreds, vc)

			}

			return cliCtx.PrintOutput(verifiablecreds)
		},
	}
}
