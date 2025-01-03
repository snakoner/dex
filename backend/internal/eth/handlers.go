package eth

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/snakoner/dex/internal/models"
)

func setHttpError(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}

/*
DexPool
*/
func (e *EthereumServer) LiquidityHandle(w http.ResponseWriter, r *http.Request) {
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

	response := &models.PairInt64Response{
		Value0: num0.Int64(),
		Value1: num1.Int64(),
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

// may be too heavy. Better calculate on front
func (e *EthereumServer) OutputAmountHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pair := vars["pair"]
	amount, err := strconv.ParseInt(vars["amount"], 10, 64)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	in, err := strconv.ParseInt(vars["in"], 10, 64)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := strconv.ParseInt(vars["out"], 10, 64)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := e.pools[pair]; !ok {
		e.logger.Error(errUnknownPair)
		setHttpError(w, errUnknownPair.Error(), http.StatusBadRequest)
		return
	}

	pool := e.pools[pair]
	outputAmount, err := pool.pool.httpInst.GetOutputAmount(&bind.CallOpts{}, big.NewInt(amount), big.NewInt(in), big.NewInt(out))
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.SingleInt64Response{
		Value: outputAmount.Int64(),
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

func (e *EthereumServer) AmountToAddHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pair := vars["pair"]

	if vars["forward"] != "true" && vars["forward"] != "false" {
		e.logger.Error(errWrongDirection)
		setHttpError(w, errWrongDirection.Error(), http.StatusBadRequest)
		return
	}
	forward := vars["forward"] == "true"

	amount, err := strconv.ParseInt(vars["amount"], 10, 64)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := e.pools[pair]; !ok {
		e.logger.Error(errUnknownPair)
		setHttpError(w, errUnknownPair.Error(), http.StatusBadRequest)
		return
	}

	pool := e.pools[pair]
	outputAmount, err := pool.pool.httpInst.GetAmountToAdd(&bind.CallOpts{}, big.NewInt(amount), forward)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.SingleInt64Response{
		Value: outputAmount.Int64(),
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

func (e *EthereumServer) AmountFromLpHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pair := vars["pair"]

	amount, err := strconv.ParseInt(vars["amount"], 10, 64)
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := e.pools[pair]; !ok {
		e.logger.Error(errUnknownPair)
		setHttpError(w, errUnknownPair.Error(), http.StatusBadRequest)
		return
	}

	pool := e.pools[pair]
	token0Amount, token1Amount, err := pool.pool.httpInst.GetAmountsFromLp(&bind.CallOpts{}, big.NewInt(amount))
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.PairInt64Response{
		Value0: token0Amount.Int64(),
		Value1: token1Amount.Int64(),
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

func (e *EthereumServer) FeeHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pair := vars["pair"]

	if _, ok := e.pools[pair]; !ok {
		e.logger.Error(errUnknownPair)
		setHttpError(w, errUnknownPair.Error(), http.StatusBadRequest)
		return
	}

	pool := e.pools[pair]
	fee, err := pool.pool.httpInst.Fee(&bind.CallOpts{})
	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.SingleInt64Response{
		Value: fee.Int64(),
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

/*
LiquidityProviderERC20
*/
func (e *EthereumServer) LpAccountBalanceHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pair := vars["pair"]
	account := vars["account"]
	if _, ok := e.pools[pair]; !ok {
		e.logger.Error(errUnknownPair)
		setHttpError(w, errUnknownPair.Error(), http.StatusBadRequest)
		return
	}

	pool := e.pools[pair]
	balance, err := pool.lp.httpInst.BalanceOf(
		&bind.CallOpts{},
		common.HexToAddress(account),
	)

	if err != nil {
		e.logger.Error(err)
		setHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &models.SingleInt64Response{
		Value: balance.Int64(),
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
