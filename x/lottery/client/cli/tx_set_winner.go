package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSetWinner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-winner [winner]",
		Short: "Broadcast message setWinner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argWinner := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetWinner(
				clientCtx.GetFromAddress().String(),
				argWinner,
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
