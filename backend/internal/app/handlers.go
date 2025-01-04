package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/snakoner/dex/internal/models"
)

func (a *App) GetPoolsHandle(w http.ResponseWriter, r *http.Request) {
	response := &models.PoolsResponse{
		FactoryAddress: a.config.FactoryAddress,
	}

	for _, p := range a.config.Pairs {
		tokenNames := strings.Split(p.Name, "-")
		if len(tokenNames) != 2 {
			a.logger.Error(errCantParsePairName)
			http.Error(w, errCantParsePairName.Error(), http.StatusInternalServerError)
			return
		}

		tokenLpAddress, err := a.ethSrv.GetLP(p.Name)
		if err != nil {
			a.logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		poolAddress, err := a.ethSrv.GetPool(p.Name)
		if err != nil {
			a.logger.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		pool := &models.PoolInfo{
			NameA:   tokenNames[0],
			NameB:   tokenNames[1],
			TokenA:  p.TokenA,
			TokenB:  p.TokenB,
			TokenLP: tokenLpAddress,
			Pool:    poolAddress,
		}

		response.Pairs = append(response.Pairs, *pool)
	}

	b, err := json.Marshal(response)
	if err != nil {
		a.logger.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
