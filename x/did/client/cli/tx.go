package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	_ "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/allinbits/cosmos-cash-poa/x/did/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	didTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	didTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding # 1
		// TODO: Add tx based commands
		// GetCmd<Action>(cdc)
	)...)

	return didTxCmd
}

// Example:
/*
// GetCmd<Action> is the CLI command for doing <Action>
func GetCmd<Action>(cdc *codec.Codec) *cobra.Command {
 	return &cobra.Command{
 		Use:   "Describe your action cmd",
 		Short: "Provide a short description on the cmd",
 		Args:  cobra.ExactArgs(2), // Does your request require arguments
 		RunE: func(cmd *cobra.Command, args []string) error {
 			cliCtx := context.NewCLIContext().WithCodec(cdc)
 			inBuf := bufio.NewReader(cmd.InOrStdin())
 			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

 			msg := types.NewMsg<Action>(Action params)
 			err = msg.ValidateBasic()
 			if err != nil {
 				return err
 			}

 			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
 		},
	}
}
*/
