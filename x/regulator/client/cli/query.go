package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/allinbits/cosmos-cash-poc/x/regulator/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group regulator queries under a subcommand
	regulatorQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	regulatorQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdRegulatorAll(queryRoute, cdc),
		)...,
	)

	return regulatorQueryCmd
}

func GetCmdRegulatorAll(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "regulators",
		Short: "regulators",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			resKVs, _, err := cliCtx.QuerySubspace(types.RegulatorKey, "regulator")
			if err != nil {
				return err
			}

			var regualtors []types.Regulator
			for _, kv := range resKVs {
				doc := types.Regulator{}
				cdc.UnmarshalBinaryBare(kv.Value, &doc)
				regualtors = append(regualtors, doc)

			}

			return cliCtx.PrintOutput(regualtors)
		},
	}
}
