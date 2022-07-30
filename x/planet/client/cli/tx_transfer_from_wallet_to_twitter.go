package cli

import (
	"strconv"

	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransferFromWalletToTwitter() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-wallet-to-twitter [username] [coin]",
		Short: "Broadcast message transferFromWalletToTwitter",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUsername := args[0]
			argCoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferFromWalletToTwitter(
				clientCtx.GetFromAddress().String(),
				argUsername,
				argCoin,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
