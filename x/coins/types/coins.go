package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Returns default gnrt coin
func GetDefaultCoin() Coin {
	// Default generator record - threshold is 2m coins
	return Coin{Name: "generator", Symbol: "gnrt", Default: true, PosminingThreshold: sdk.NewIntWithDecimal(2000000, 6)}
}

// Returns a stub record of the coin
func GetCoinStub(symbol string) Coin {
	if symbol == "gnrt" {
		return GetDefaultCoin()
	}

	return Coin{ Symbol: symbol, Default: false, PosminingThreshold: sdk.NewInt(0)}
}

// Returns default sdk coins
func GetDefaultCoins(amnt sdk.Int) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin("gnrt", amnt))
}

func GetGenesisWallet() sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32("gnrt1wjc2pg9e3p8sdll4kg9ssj44nhm7ce4prapsxf")

	return addr
}
