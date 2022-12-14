package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/karthik340/game/testutil/network"
	"github.com/karthik340/game/testutil/nullify"
	"github.com/karthik340/game/x/lottery/client/cli"
    "github.com/karthik340/game/x/lottery/types"
)

func networkWithValidatorsWinnerObjects(t *testing.T) (*network.Network, types.ValidatorsWinner) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
    require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	validatorsWinner := &types.ValidatorsWinner{}
	nullify.Fill(&validatorsWinner)
	state.ValidatorsWinner = validatorsWinner
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.ValidatorsWinner
}

func TestShowValidatorsWinner(t *testing.T) {
	net, obj := networkWithValidatorsWinnerObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.ValidatorsWinner
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowValidatorsWinner(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetValidatorsWinnerResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ValidatorsWinner)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ValidatorsWinner),
				)
			}
		})
	}
}

