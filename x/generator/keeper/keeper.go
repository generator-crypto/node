package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/generator-crypto/node/x/coins"
	"github.com/generator-crypto/node/x/emission"
	"github.com/generator-crypto/node/x/posmining"
	"github.com/generator-crypto/node/x/structure"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/generator-crypto/node/x/generator/types"
)

// Keeper of the generator store
type Keeper struct {
	accountKeeper    auth.AccountKeeper
	coinKeeper       bank.Keeper
	CoinsKeeper      coins.Keeper
	structureKeeper  structure.Keeper
	posminingKeeper posmining.Keeper
	emissionKeeper   emission.Keeper
	supplyKeeper     supply.Keeper
	slashingKeeper   slashing.Keeper

	cdc *codec.Codec
}

// NewKeeper creates a generator keeper
func NewKeeper(cdc *codec.Codec, accountKeeper auth.AccountKeeper, coinKeeper bank.Keeper, structureKeeper structure.Keeper, posminingKeeper posmining.Keeper, emissionKeeper emission.Keeper, supplyKeeper supply.Keeper, slashingKeeper slashing.Keeper, coinsKeeper coins.Keeper) Keeper {
	return Keeper{
		cdc:              cdc,
		accountKeeper:    accountKeeper,
		coinKeeper:       coinKeeper,
		structureKeeper:  structureKeeper,
		posminingKeeper: posminingKeeper,
		emissionKeeper:   emissionKeeper,
		supplyKeeper:     supplyKeeper,
		slashingKeeper:   slashingKeeper,
		CoinsKeeper:      coinsKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
