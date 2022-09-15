package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/spf13/cobra"
)

func CmdListStoredMiner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stored-miner",
		Short: "list all storedMiner",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStoredMinerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StoredMinerAll(context.Background(), params)
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

func CmdShowStoredMiner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stored-miner [index]",
		Short: "shows a storedMiner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetStoredMinerRequest{
				Index: argIndex,
			}

			res, err := queryClient.StoredMiner(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
