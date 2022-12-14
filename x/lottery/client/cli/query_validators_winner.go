package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/karthik340/game/x/lottery/types"
	"github.com/spf13/cobra"
)

func CmdShowValidatorsWinner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validators-winner",
		Short: "shows validatorsWinner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetValidatorsWinnerRequest{
				Validator: args[0],
			}

			res, err := queryClient.ValidatorsWinner(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
