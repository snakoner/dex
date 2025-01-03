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
