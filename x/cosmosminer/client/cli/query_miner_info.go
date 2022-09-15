package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/spf13/cobra"
)

func CmdShowMinerInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-miner-info",
		Short: "shows MinerInfo",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetMinerInfoRequest{}

			res, err := queryClient.MinerInfo(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
