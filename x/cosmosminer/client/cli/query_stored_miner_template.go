package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/spf13/cobra"
)

func CmdListStoredMinerTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stored-miner-template",
		Short: "list all storedMinerTemplate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStoredMinerTemplateRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StoredMinerTemplateAll(context.Background(), params)
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

func CmdShowStoredMinerTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stored-miner-template [index]",
		Short: "shows a storedMinerTemplate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetStoredMinerTemplateRequest{
				Index: argIndex,
			}

			res, err := queryClient.StoredMinerTemplate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
