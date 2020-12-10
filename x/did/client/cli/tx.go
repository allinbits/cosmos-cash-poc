package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/allinbits/cosmos-cash-poa/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/viper"
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
		GetCmdCreateDidDocument(cdc),
		GetCmdCreateVerifiableCredential(cdc),
	)...)

	return didTxCmd
}

// GetCmdCreateDidDocument is the CLI command for sending a CreateDidDocument transaction
func GetCmdCreateDidDocument(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-did-document",
		Short: "create an did document for an address",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()
			id := types.DidIdentifer + accAddr.String()
			keybase, err := keys.NewKeyring(sdk.KeyringServiceName(),
				viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), inBuf)
			info, err := keybase.GetByAddress(accAddr)
			if err != nil {
				return err
			}
			pubKeyBase58 := info.GetPubKey().Address()
			pubKey := types.NewPubKey(id, "Ed25519VerificationKey2018", accAddr, pubKeyBase58.String())
			authentication := types.PubKeys{pubKey}

			msg := types.NewMsgCreateDidDocument(types.Context, id, authentication, nil, accAddr)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdCreateDidDocument is the CLI command for sending a CreateDidDocument transaction
func GetCmdCreateVerifiableCredential(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-verifiable-credential [address]",
		Short: "create an verifiable cred for an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			// if err := cliCtx.EnsureAccountExists(); err != nil {
			// 	return err
			// }

			accAddr := cliCtx.GetFromAddress()
			id := types.DidIdentifer + accAddr.String()

			msg := types.NewMsgCreateVerifiableCredential(types.DidIdentifer+args[0], types.VcContext, id, "VerifiableCredential", accAddr.String(), types.Proof{}, accAddr)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
