package cli

import (
	"strconv"

	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdTransferFromTwitterToTwitterByMerces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-twitter-to-twitter-by-merces [from-username] [to-username] [denom] [amount]",
		Short: "Broadcast message transferFromTwitterToTwitterByMerces",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFromUsername := args[0]
			argToUsername := args[1]
			argDenom := args[2]
			argAmount, err := cast.ToInt64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferFromTwitterToTwitterByMerces(
				clientCtx.GetFromAddress().String(),
				argFromUsername,
				argToUsername,
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
