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
		pool := &models.PoolInfo{
			NameA:  tokenNames[0],
			NameB:  tokenNames[1],
			TokenA: p.TokenA,
			TokenB: p.TokenB,
			Pool:   p.Pool,
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
