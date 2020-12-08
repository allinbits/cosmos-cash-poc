package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash-poa/x/regulator/types"
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
	regulatorTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	regulatorTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateRegulator(cdc),
	)...)

	return regulatorTxCmd
}

// GetCmdCreateRegulator is the CLI command for sending a CreateDidDocument transaction
func GetCmdCreateRegulator(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-regulator [address]",
		Short: "create a regulator for an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()
			regAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateRegualtor(regAddr, accAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
