package eth

import (
	"net/http"
)

func setHttpError(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}

// todo: filter by indexed round

// RoundHandler is an HTTP handler that processes requests to retrieve bid events
// for a specific lottery round. It accepts a round number from the URL, filters
// the bid events from the Ethereum contract based on the provided round number,
// and returns the corresponding events in the response.
//
// Parameters:
//   - w: The HTTP response writer used to send responses back to the client.
//   - r: The HTTP request containing the round number in the URL.
//
// Response:
//   - A JSON object containing the list of bid events for the specified round if successful.
//   - A corresponding HTTP error status if there is an issue (e.g., 400 for bad requests or 500 for server errors).
func (e *EthereumServer) RoundHandler(w http.ResponseWriter, r *http.Request) {

}

// WinnerHandler is an HTTP handler that processes requests to retrieve information
// about the selected winner for a specific lottery round. It accepts the round number
// from the URL parameter, filters the winner selection events from the Ethereum contract,
// and returns the selected winner's details in the response.
//
// Parameters:
//   - w: The HTTP response writer used to send responses back to the client.
//   - r: The HTTP request containing the round number in the URL.
//
// Response:
//   - A JSON object containing the winner's details if successful.
//   - A corresponding HTTP error status if there is an issue (e.g., 400 for bad requests or 500 for server errors).
func (e *EthereumServer) WinnerHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *EthereumServer) AllTimeRewardHander(w http.ResponseWriter, r *http.Request) {

}
