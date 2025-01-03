package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Pair struct {
	Name   string
	TokenA string
	TokenB string
}

type Config struct {
	Name string `json:"name"`
	Log  struct {
		Level  string `json:"level"`
		Output string `json:"output"`
	} `json:"log"`
	Service struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"service"`
	Infura struct {
		HttpProvider string `json:"http-provider"`
		WssProvider  string `json:"wss-provider"`
	} `json:"infura"`
	Private struct {
		Evm      string
		Provider string
	}
	FactoryAddress string `json:"factory-address"`
	Pairs          []*Pair
}

type pairsJSON struct {
	Pairs []struct {
		NameA  string `json:"name-a"`
		NameB  string `json:"name-b"`
		TokenA string `json:"token-a"`
		TokenB string `json:"token-b"`
	} `json:"pairs"`
}

/*
Purpose of this function to swap name of pair,
if tokenA address is bigger than tokenB address
for subsequent receipt of an address from factory
*/
func createPairs(pairs *pairsJSON, config *Config) {
	for _, p := range pairs.Pairs {
		var nameA, nameB string
		var tokenA, tokenB string
		if p.TokenB < p.TokenA {
			nameA, nameB = p.NameB, p.NameA
			tokenA, tokenB = p.TokenB, p.TokenA
		} else {
			nameA, nameB = p.NameA, p.NameB
			tokenA, tokenB = p.TokenA, p.TokenB
		}

		newPair := &Pair{
			Name:   fmt.Sprintf("%s-%s", nameA, nameB),
			TokenA: tokenA,
			TokenB: tokenB,
		}

		config.Pairs = append(config.Pairs, newPair)
	}
}

func New(configPath string) (*Config, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := json.Unmarshal(b, config); err != nil {
		return nil, err
	}

	pairsJSON := &pairsJSON{}
	if err := json.Unmarshal(b, pairsJSON); err != nil {
		return nil, err
	}

	createPairs(pairsJSON, config)

	if err := setPrivate(config); err != nil {
		return nil, err
	}

	return config, nil
}

func setPrivate(config *Config) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	val, ok := os.LookupEnv("EVM_PRIVATE_KEY")
	if !ok {
		return errors.New("cant get evm private key")
	}

	config.Private.Evm = val

	val, ok = os.LookupEnv("ALCHEMY_API_KEY")
	if !ok {
		return errors.New("cant get provider api key")
	}

	config.Private.Provider = val

	return nil
}
