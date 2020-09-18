package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	issuerTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s issuer subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	issuerTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateIssuer(cdc),
	)...)

	return issuerTxCmd
}

// GetCmdCreateValidator is the CLI command for sending a CreateValidator transaction
func GetCmdCreateIssuer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-issuer [name] [address] [token] [amount]",
		Short: "create an issuer and a token pair",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()

			issuerAccount, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateIssuer(args[0], issuerAccount, args[2], args[3], accAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
