package eth

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/mux"
	"github.com/snakoner/dex/internal/models"
)

func setHttpError(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}

func (e *EthereumServer) LiquidityHandler(w http.ResponseWriter, r *http.Request) {
	pair := mux.Vars(r)["pair"]
	if _, ok := e.pools[pair]; !ok {
		e.logger.Error(errUnknownPair)
		setHttpError(w, errUnknownPair.Error(), http.StatusBadRequest)
		return
	}

	pool := e.pools[pair]
	num0, err := pool.pool.httpInst.GetReserve0(&bind.CallOpts{})
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	num1, err := pool.pool.httpInst.GetReserve1(&bind.CallOpts{})
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.LiquidityResponse{
		NumberA: num0.Int64(),
		NumberB: num1.Int64(),
	}

	b, err := json.Marshal(response)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
