package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Pair struct {
	Name   string `json:"name"`
	TokenA string `json:"token-a"`
	TokenB string `json:"token-b"`
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
	FactoryAddress string  `json:"factory-address"`
	Pairs          []*Pair `json:"pairs"`
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
