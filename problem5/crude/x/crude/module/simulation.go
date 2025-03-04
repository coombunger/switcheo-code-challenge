package crude

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"crude/testutil/sample"
	crudesimulation "crude/x/crude/simulation"
	"crude/x/crude/types"
)

// avoid unused import issue
var (
	_ = crudesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreate = "op_weight_msg_create"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreate int = 100

	opWeightMsgUpdate = "op_weight_msg_update"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdate int = 100

	opWeightMsgDelete = "op_weight_msg_delete"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDelete int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	crudeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&crudeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreate int
	simState.AppParams.GetOrGenerate(opWeightMsgCreate, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCreate = defaultWeightMsgCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreate,
		crudesimulation.SimulateMsgCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdate int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdate, &weightMsgUpdate, nil,
		func(_ *rand.Rand) {
			weightMsgUpdate = defaultWeightMsgUpdate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdate,
		crudesimulation.SimulateMsgUpdate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDelete int
	simState.AppParams.GetOrGenerate(opWeightMsgDelete, &weightMsgDelete, nil,
		func(_ *rand.Rand) {
			weightMsgDelete = defaultWeightMsgDelete
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDelete,
		crudesimulation.SimulateMsgDelete(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreate,
			defaultWeightMsgCreate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudesimulation.SimulateMsgCreate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdate,
			defaultWeightMsgUpdate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudesimulation.SimulateMsgUpdate(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDelete,
			defaultWeightMsgDelete,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudesimulation.SimulateMsgDelete(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
