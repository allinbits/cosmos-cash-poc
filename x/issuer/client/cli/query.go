package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	//sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group issuer queries under a subcommand
	issuerQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	issuerQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdIssuersAll(queryRoute, cdc),
		)...,
	)

	return issuerQueryCmd
}

func GetCmdIssuersAll(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "issuers",
		Short: "issuers",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.IssuersKey, queryRoute)
			if err != nil {
				return err
			}

			var validators []types.Issuer
			for _, kv := range resKVs {
				validator := types.Issuer{}
				cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &validator)
				validators = append(validators, validator)

			}

			return cliCtx.PrintOutput(validators)
		},
	}
}
