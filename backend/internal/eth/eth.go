package eth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/snakoner/dex/internal/bindings/factory"
	"github.com/snakoner/dex/internal/bindings/liquidityprovider"
	"github.com/snakoner/dex/internal/bindings/pool"
	"github.com/snakoner/dex/internal/config"
)

type Signer struct {
	privateKey *ecdsa.PrivateKey
}

type ExchangeObject struct {
	lp struct {
		address  common.Address
		httpCli  *ethclient.Client
		httpInst *liquidityprovider.Lp
		wsCli    *ethclient.Client
		wsInst   *liquidityprovider.Lp
	}
	pool struct {
		address  common.Address
		httpCli  *ethclient.Client
		httpInst *pool.Pool
		wsCli    *ethclient.Client
		wsInst   *pool.Pool
	}
}

type Factory struct {
	address  common.Address
	httpCli  *ethclient.Client
	httpInst *factory.Factory
	wsCli    *ethclient.Client
	wsInst   *factory.Factory
}

type EthereumServer struct {
	exhanges map[string]ExchangeObject
	factory  *Factory
	logger   *logrus.Logger
	signer   *Signer
}

func New(config *config.Config) (*EthereumServer, error) {
	return nil, nil
}

func (e *EthereumServer) Stop(name string) {
	exchange, ok := e.exhanges[name]
	if ok {
		clients := []*ethclient.Client{
			exchange.lp.httpCli,
			exchange.lp.wsCli,
			exchange.pool.httpCli,
			exchange.pool.wsCli,
		}

		for _, cli := range clients {
			if cli != nil {
				cli.Close()
			}
		}
	}
}
