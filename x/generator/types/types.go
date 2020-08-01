package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	GNRT                 = "gnrt"
	POINTS               = 6        // points after comma
	INITIAL              = 10000000 // initial emission
	PARAMINING_THRESHOLD = 2000000  // posmining threshold
)

func NewgnrtCoin(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(GNRT, amount)
}

func NewCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(GNRT, amount)
}
func NewCoins(amount sdk.Int) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(GNRT, amount))
}

// Since we cannot make constans as objects
func GetMaxLevel() sdk.Int {
	return sdk.NewInt(100)
}

func GetPosminingThreshold() sdk.Int {
	return sdk.NewIntWithDecimal(PARAMINING_THRESHOLD, POINTS)
}
