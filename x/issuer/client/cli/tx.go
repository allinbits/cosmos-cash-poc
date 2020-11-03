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
		GetCmdMintToken(cdc),
		GetCmdBurnToken(cdc),
		GetCmdFreezeToken(cdc),
		GetCmdUnfreezeToken(cdc),
		GetCmdWithdrawToken(cdc),
		GetCmdFreezeAccount(cdc),
	)...)

	return issuerTxCmd
}

// GetCmdCreateIssuer is the CLI command for sending a CreateIssuer transaction
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

// GetCmdMintToken is the CLI command for sending a MintToken transaction
func GetCmdMintToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "mint-token [token] [amount]",
		Short: "mint tokens",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgMintToken(args[0], args[1], accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdBurnToken is the CLI command for sending a BurnToken transaction
func GetCmdBurnToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "burn-token [token] [amount]",
		Short: "burn token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgBurnToken(args[0], args[1], accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdFreezeToken is the CLI command for sending a FreezeToken transaction
func GetCmdFreezeToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "freeze-token [token]",
		Short: "freeze token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgFreezeToken(args[0], accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUnfreezeToken is the CLI command for sending a UnfreezeToken transaction
func GetCmdUnfreezeToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "unfreeze-token [token]",
		Short: "unfreeze token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgUnfreezeToken(args[0], accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdWithdrawToken is the CLI command for sending a WithdrawToken transaction
func GetCmdWithdrawToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-token [token] [amount]",
		Short: "withdraw token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgWithdrawToken(args[0], args[1], accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdUnfreezeToken is the CLI command for sending a UnfreezeToken transaction
func GetCmdFreezeAccount(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "freeze-account [account]",
		Short: "freeze account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			freezeAccAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			accAddr := cliCtx.GetFromAddress()

			msg := types.NewMsgFreezeAccount(freezeAccAddr, accAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
