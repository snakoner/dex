package eth

import (
	"net/http"
)

func setHttpError(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}

func (e *EthereumServer) RoundHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *EthereumServer) WinnerHandler(w http.ResponseWriter, r *http.Request) {

}

func (e *EthereumServer) AllTimeRewardHander(w http.ResponseWriter, r *http.Request) {

}
