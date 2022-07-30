package cli

import (
	"context"

	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListTwitterCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-twitter-coins",
		Short: "list all twitterCoins",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			username := args[0]

			params := &types.QueryAllTwitterCoinsRequest{
				Username:   username,
				Pagination: pageReq,
			}

			res, err := queryClient.TwitterCoinsAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowTwitterCoins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-twitter-coins [username] [index]",
		Short: "shows a twitterCoins",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			username := args[0]
			argIndex := args[1]

			params := &types.QueryGetTwitterCoinsRequest{
				Username: username,
				Index:    argIndex,
			}

			res, err := queryClient.TwitterCoins(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
