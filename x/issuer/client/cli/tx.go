package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	issuerTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	issuerTxCmd.AddCommand(flags.PostCommands(
	// this line is used by starport scaffolding
	)...)

	return issuerTxCmd
}
