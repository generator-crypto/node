package posmining

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/generator-crypto/node/x/coins"
	"github.com/generator-crypto/node/x/posmining/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	defaultCoin := coins.GetDefaultCoin()

	for _, record := range data.Records {
		k.SetPosmining(ctx, record, defaultCoin)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	var records []types.Posmining

	iterator := k.GetPosminingIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		var posmining types.Posmining

		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &posmining)

		records = append(records, posmining)
	}

	return NewGenesisState(ctx, records)
}
