package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/luktech-labs/cosmos-miner/x/cosmosminer/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cosmosminer queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowMinerInfo())
	cmd.AddCommand(CmdListStoredMinerTemplate())
	cmd.AddCommand(CmdShowStoredMinerTemplate())
	cmd.AddCommand(CmdListStoredMiner())
	cmd.AddCommand(CmdShowStoredMiner())
	cmd.AddCommand(CmdListMinerBalance())
	cmd.AddCommand(CmdShowMinerBalance())
	// this line is used by starport scaffolding # 1

	return cmd
}
