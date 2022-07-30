package cli

import (
	"context"

	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListDNSRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-dns-registry",
		Short: "list all DNSRegistry",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDNSRegistryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DNSRegistryAll(context.Background(), params)
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

func CmdShowDNSRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-dns-registry [domain]",
		Short: "shows a DNSRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDomain := args[0]

			params := &types.QueryGetDNSRegistryRequest{
				Domain: argDomain,
			}

			res, err := queryClient.DNSRegistry(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
