package lottery

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/karthik340/game/testutil/sample"
	lotterysimulation "github.com/karthik340/game/x/lottery/simulation"
	"github.com/karthik340/game/x/lottery/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = lotterysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgPlaceBet = "op_weight_msg_place_bet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlaceBet int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	lotteryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&lotteryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	lotteryParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinFee), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(lotteryParams.MinFee))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinBet), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(lotteryParams.MinBet))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxBet), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(lotteryParams.MaxBet))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinTxn), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(lotteryParams.MinTxn))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgPlaceBet int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPlaceBet, &weightMsgPlaceBet, nil,
		func(_ *rand.Rand) {
			weightMsgPlaceBet = defaultWeightMsgPlaceBet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlaceBet,
		lotterysimulation.SimulateMsgPlaceBet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
