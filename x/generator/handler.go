package generator

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/generator-crypto/node/x/bank"
	coinTypes "github.com/generator-crypto/node/x/coins/types"
	"github.com/generator-crypto/node/x/emission"
	"github.com/generator-crypto/node/x/generator/types"
	"time"
)

type RestoreUnbonding struct {
	Addr  string
	Coins int64
}

var Restores = []RestoreUnbonding{
	{"gnrt1lqplawlvt48aaksay72qxk043cqvg88j8djt25", 60000},
	{"gnrt1a05x80vtl9p2vk9m58dka390vg3aslxtkgye2n", 3000},
	{"gnrt1ursvfauwc2gvjps7myepuw0p6wrv20eqwv7t2r", 3586},
	{"gnrt1cxepp96w3ea5ghvx94h689yxrade2cj3sc8vdr", 383402},
	{"gnrt1j3c9s6dkw98h2x56frstlfah48qkeqk04tg2f5", 80},
	{"gnrt1js8wtk8qdqa8eja8qgfnk7vy7w3vhpszr2fl44", 2},
	{"gnrt1z52hp00w34vz2dhan9vacyc3lezjetfajvja8l", 100},
	{"gnrt1z3zv4cja9wt5htwrh2h0l6ar03eh0q3u3j5crq", 126},
}

// NewHandler creates an sdk.Handler for all the generator type messages
func NewHandler(k Keeper, paramsKeeper params.Keeper, bankKeeper bank.Keeper, emissionKeeper emission.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgChangeParams:
			return handleChangeParams(ctx, k, msg, paramsKeeper)
		case types.MsgRestoreUnbonding:
			return handleRestoreUnbonding(ctx, k, msg, bankKeeper, emissionKeeper)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleChangeParams(ctx sdk.Context, k Keeper, msg types.MsgChangeParams, paramsKeeper params.Keeper) (*sdk.Result, error) {
	if !msg.Owner.Equals(types.GetGenesisWallet()) {
		return nil, sdkerrors.Wrapf(params.ErrSettingParameter, "only genesis can call this method")
	}

	ss, ok := paramsKeeper.GetSubspace("staking")

	if !ok {
		return nil, sdkerrors.Wrap(params.ErrUnknownSubspace, "staking")
	}

	var NewValue time.Duration = time.Hour * 24 * 3

	bin, _ := codec.Cdc.MarshalJSON(NewValue)

	if err := ss.Update(ctx, []byte("UnbondingTime"), bin); err != nil {
		fmt.Println(err)

		return nil, sdkerrors.Wrapf(params.ErrSettingParameter, "key: %s, value: %s, err: %s", "unbonding_time", "", err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleRestoreUnbonding(ctx sdk.Context, k Keeper, msg types.MsgRestoreUnbonding, bankKeeper bank.Keeper, emissionKeeper emission.Keeper) (*sdk.Result, error) {
	if !msg.Owner.Equals(types.GetGenesisWallet()) {
		return nil, sdkerrors.Wrapf(params.ErrSettingParameter, "only genesis can call this method")
	}

	for _, v := range Restores {
		sdkAddr, err := sdk.AccAddressFromBech32(v.Addr)

		if err != nil {
			panic(err)
		}

		coinsAmount := sdk.NewIntWithDecimal(v.Coins, 6)

		gnrt := coinTypes.GetDefaultCoin()

		coins := sdk.NewCoins(sdk.NewCoin(gnrt.Symbol, coinsAmount))

		_, err = bankKeeper.AddCoins(ctx, sdkAddr, coins)

		if err != nil {
			panic(err)
		}

		emissionKeeper.Add(ctx, coinsAmount, gnrt)
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
