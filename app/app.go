package app

import (
	"encoding/json"
	"io"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/allinbits/cosmos-cash-poa/x/did"
	didkeeper "github.com/allinbits/cosmos-cash-poa/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash-poa/x/did/types"

	"github.com/allinbits/cosmos-cash-poa/x/regulator"
	regulatorkeeper "github.com/allinbits/cosmos-cash-poa/x/regulator/keeper"
	regulatortypes "github.com/allinbits/cosmos-cash-poa/x/regulator/types"

	"github.com/allinbits/cosmos-cash-poa/x/issuer"
	issuerkeeper "github.com/allinbits/cosmos-cash-poa/x/issuer/keeper"
	issuertypes "github.com/allinbits/cosmos-cash-poa/x/issuer/types"

	"github.com/allinbits/modules/poa"
	poakeeper "github.com/allinbits/modules/poa/keeper"
	poatypes "github.com/allinbits/modules/poa/types"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

const appName = "cosmos-cash"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.poacli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.poad")
	ModuleBasics    = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		supply.AppModuleBasic{},
		poa.AppModuleBasic{},
		issuer.AppModuleBasic{},
		did.AppModuleBasic{},
		regulator.AppModuleBasic{},
	)

	maccPerms = map[string][]string{
		auth.FeeCollectorName:  nil,
		issuertypes.ModuleName: nil,
	}
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc.Seal()
}

type NewApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	keys  map[string]*sdk.KVStoreKey
	tKeys map[string]*sdk.TransientStoreKey

	subspaces map[string]params.Subspace

	accountKeeper   auth.AccountKeeper
	bankKeeper      bank.Keeper
	supplyKeeper    supply.Keeper
	paramsKeeper    params.Keeper
	poaKeeper       poakeeper.Keeper
	issuerKeeper    issuerkeeper.Keeper
	didKeeper       didkeeper.Keeper
	regulatorKeeper regulatorkeeper.Keeper
	mm              *module.Manager

	sm *module.SimulationManager
}

var _ simapp.App = (*NewApp)(nil)

func NewInitApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) *NewApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey,
		auth.StoreKey,
		supply.StoreKey,
		params.StoreKey,
		poatypes.StoreKey,
		issuertypes.StoreKey,
		didtypes.StoreKey,
		regulatortypes.StoreKey,
	)

	tKeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	var app = &NewApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
		subspaces:      make(map[string]params.Subspace),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tKeys[params.TStoreKey])
	app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[poatypes.ModuleName] = app.paramsKeeper.Subspace(poakeeper.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount,
	)

	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.subspaces[bank.ModuleName],
		app.ModuleAccountAddrs(),
	)

	app.supplyKeeper = supply.NewKeeper(
		app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		maccPerms,
	)

	app.poaKeeper = poakeeper.NewKeeper(
		app.bankKeeper,
		app.cdc,
		keys[poatypes.StoreKey],
		app.subspaces[poatypes.ModuleName],
	)

	app.issuerKeeper = issuerkeeper.NewKeeper(
		app.bankKeeper,
		app.supplyKeeper,
		app.cdc,
		keys[issuertypes.StoreKey],
	)

	app.didKeeper = didkeeper.NewKeeper(
		app.cdc,
		keys[didtypes.StoreKey],
	)

	app.regulatorKeeper = regulatorkeeper.NewKeeper(
		app.cdc,
		keys[regulatortypes.StoreKey],
	)

	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.poaKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		poa.NewAppModule(app.poaKeeper, app.bankKeeper),
		issuer.NewAppModule(app.issuerKeeper, app.bankKeeper),
		did.NewAppModule(app.didKeeper),
		regulator.NewAppModule(app.regulatorKeeper, app.didKeeper),
	)

	app.mm.SetOrderEndBlockers(poatypes.ModuleName)

	genutil.ModuleCdc = app.cdc

	app.mm.SetOrderInitGenesis(
		auth.ModuleName,
		bank.ModuleName,
		poatypes.ModuleName,
		supply.ModuleName,
		genutil.ModuleName,
		//regulatortypes.ModuleName,
	)

	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	app.SetAnteHandler(
		NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			app.issuerKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

func (app *NewApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState

	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	var authGenesisState authtypes.GenesisState
	app.cdc.MustUnmarshalJSON(genesisState["auth"], &authGenesisState)
	regulator.InitGenesis(ctx, app.regulatorKeeper, authGenesisState)

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *NewApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *NewApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *NewApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

func (app *NewApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (app *NewApp) Codec() *codec.Codec {
	return app.cdc
}

func (app *NewApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

func GetMaccPerms() map[string][]string {
	modAccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		modAccPerms[k] = v
	}
	return modAccPerms
}
