package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateMinerTemplate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-miner-template [name] [price] [prod-per-second] [min-claim-time]",
		Short: "Broadcast message createMinerTemplate",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argPrice, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argProdPerSecond, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argMinClaimTime, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMinerTemplate(
				clientCtx.GetFromAddress().String(),
				argName,
				argPrice,
				argProdPerSecond,
				argMinClaimTime,
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
