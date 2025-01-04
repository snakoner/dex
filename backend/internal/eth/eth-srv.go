package eth

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/snakoner/dex/internal/bindings/factory"
	liquidityprovider "github.com/snakoner/dex/internal/bindings/liquidity-provider"
	"github.com/snakoner/dex/internal/bindings/pool"
	"github.com/snakoner/dex/internal/config"
)

type Signer struct {
	privateKey *ecdsa.PrivateKey
}

type PoolObject struct {
	lp struct {
		address  common.Address
		httpInst *liquidityprovider.Liquidityprovider
		wsInst   *liquidityprovider.Liquidityprovider
	}
	pool struct {
		address  common.Address
		httpInst *pool.Pool
		wsInst   *pool.Pool
	}
}

type Factory struct {
	address  common.Address
	httpInst *factory.Factory
	wsInst   *factory.Factory
}

type EthereumServer struct {
	pools   map[string]PoolObject
	httpCli *ethclient.Client
	wsCli   *ethclient.Client
	factory *Factory
	logger  *logrus.Logger
	signer  *Signer
}

func (e *EthereumServer) GetLP(pair string) (string, error) {
	pool, ok := e.pools[pair]
	if !ok {
		return "", errUnknownPair
	}

	return pool.lp.address.Hex(), nil
}

func (e *EthereumServer) GetPool(pair string) (string, error) {
	pool, ok := e.pools[pair]
	if !ok {
		return "", errUnknownPair
	}

	return pool.pool.address.Hex(), nil
}

func (e *EthereumServer) setupFactory(config *config.Config) error {
	address := common.HexToAddress(config.FactoryAddress)

	// http
	httpCli, err := ethclient.Dial(fmt.Sprintf("%s%s", config.Infura.HttpProvider, config.Private.Provider))
	if err != nil {
		return err
	}

	httpInst, err := factory.NewFactory(address, httpCli)
	if err != nil {
		return err
	}

	// ws
	wssCli, err := ethclient.Dial(fmt.Sprintf("%s%s", config.Infura.WssProvider, config.Private.Provider))
	if err != nil {
		return err
	}

	wssInst, err := factory.NewFactory(address, wssCli)
	if err != nil {
		return err
	}

	e.factory = &Factory{
		address:  address,
		httpInst: httpInst,
		wsInst:   wssInst,
	}

	return nil
}

func (e *EthereumServer) setupPools(config *config.Config) error {
	for _, pair := range config.Pairs {
		name := pair.Name
		tokenA := common.HexToAddress(pair.TokenA)
		tokenB := common.HexToAddress(pair.TokenB)

		poolAddress, err := e.factory.httpInst.GetPool(&bind.CallOpts{}, tokenA, tokenB)

		if err != nil {
			return err
		}

		// http
		poolHttpInst, err := pool.NewPool(poolAddress, e.httpCli)
		if err != nil {
			return err
		}

		liquidityProviderAddress, err := poolHttpInst.LpToken(&bind.CallOpts{})
		if err != nil {
			return err
		}
		liquidityProviderHttpInst, err := liquidityprovider.NewLiquidityprovider(liquidityProviderAddress, e.httpCli)
		if err != nil {
			return err
		}

		// ws
		poolWsInst, err := pool.NewPool(poolAddress, e.wsCli)
		if err != nil {
			return err
		}

		liquidityProviderWsInst, err := liquidityprovider.NewLiquidityprovider(liquidityProviderAddress, e.wsCli)
		if err != nil {
			return err
		}

		obj := &PoolObject{
			lp: struct {
				address  common.Address
				httpInst *liquidityprovider.Liquidityprovider
				wsInst   *liquidityprovider.Liquidityprovider
			}{
				address:  liquidityProviderAddress,
				httpInst: liquidityProviderHttpInst,
				wsInst:   liquidityProviderWsInst,
			},
			pool: struct {
				address  common.Address
				httpInst *pool.Pool
				wsInst   *pool.Pool
			}{
				address:  poolAddress,
				httpInst: poolHttpInst,
				wsInst:   poolWsInst,
			},
		}

		e.pools[name] = *obj
	}

	return nil
}

func New(config *config.Config) (*EthereumServer, error) {
	privateKey, err := crypto.HexToECDSA(config.Private.Evm)
	if err != nil {
		return nil, err
	}

	httpCli, err := ethclient.Dial(fmt.Sprintf("%s%s", config.Infura.HttpProvider, config.Private.Provider))
	if err != nil {
		return nil, err
	}

	wsCli, err := ethclient.Dial(fmt.Sprintf("%s%s", config.Infura.WssProvider, config.Private.Provider))
	if err != nil {
		return nil, err
	}

	e := &EthereumServer{
		logger: logrus.New(),
		signer: &Signer{
			privateKey: privateKey,
		},
		httpCli: httpCli,
		wsCli:   wsCli,
		pools:   make(map[string]PoolObject),
	}

	// setup factory
	if err := e.setupFactory(config); err != nil {
		return nil, err
	}

	// setup pools
	if err := e.setupPools(config); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *EthereumServer) Stop() {
	if e.httpCli != nil {
		e.httpCli.Close()
	}

	if e.wsCli != nil {
		e.wsCli.Close()
	}
}
