package cli

import (
	"strconv"

	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransferFromTwitterToWalletByMerces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-twitter-to-wallet-by-merces [username] [address] [denom] [amount]",
		Short: "Broadcast message transferFromTwitterToWalletByMerces",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUsername := args[0]
			argAddress := args[1]
			argDenom := args[2]
			argAmount, err := cast.ToInt64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferFromTwitterToWalletByMerces(
				clientCtx.GetFromAddress().String(),
				argUsername,
				argAddress,
				argDenom,
				argAmount,
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
