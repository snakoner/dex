package models

type BidEvent struct {
	Account   string `json:"account"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Round     int64  `json:"round"`
}

type BidResponse struct {
	Events []*BidEvent `json:"events"`
}

type WinnerSelectedEvent struct {
	Account       string `json:"account"`
	Amount        int64  `json:"amount"`
	Round         int64  `json:"round"`
	Timestamp     int64  `json:"timestamp"`
	RoundFinished bool   `json:"round-finished"`
}

type WinnerSelectedResponse struct {
	Winner WinnerSelectedEvent `json:"winner"`
}

// eth value
type AllTimeRewardResponse struct {
	Reward string `json:"reward"`
}
