package models

type LiquidityResponse struct {
	NumberA int64 `json:"number-a"`
	NumberB int64 `json:"number-b"`
}

// getOutputAmount, getAmountToAdd
type OutputAmountResponse struct {
	Amount int64 `json:"amount"`
}

// getOutputAmount, getAmountToAdd
type AmountFromLpResponse struct {
	NumberA int64 `json:"number-a"`
	NumberB int64 `json:"number-b"`
}

type SingleInt64Response struct {
	Value int64 `json:"value"`
}

type PairInt64Response struct {
	Value0 int64 `json:"value0"`
	Value1 int64 `json:"value1"`
}
