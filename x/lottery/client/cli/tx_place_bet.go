package cli

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPlaceBet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "place-bet [fee] [bet]",
		Short: "Broadcast message placeBet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fee, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			bet, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPlaceBet(
				clientCtx.GetFromAddress().String(),
				fee,
				bet,
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
