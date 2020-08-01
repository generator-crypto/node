package bank

import (
	"github.com/generator-crypto/node/x/bank/keeper"
)

var (
	NewKeeper = keeper.NewKeeper
)

type (
	Keeper = keeper.Keeper
)
