package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var DefaultConfig = Config{
	DB: DBConfig{"sqlite3", ":memory:"},
	Code: CodeConfig{
		Alphabet: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		Size:     4,
	},
	ListenAddress: "0.0.0.0:8080",
}

type DBConfig struct {
	Dialect string `json:"dialect"`
	DSN     string `json:"dsn"`
}
type CodeConfig struct {
	Alphabet string `json:"alphabet"`
	Size     int    `json:"size"`
}

type Config struct {
	DB            DBConfig   `json:"db"`
	Code          CodeConfig `json:"code"`
	ListenAddress string     `json:"listen_address"`
}

func LoadConfigFromFileOrDefault(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		copyDefaultConfig := DefaultConfig
		return &copyDefaultConfig, nil
	}

	configData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read config failed")
	}

	var c Config
	if err := json.Unmarshal(configData, &c); err != nil {
		return nil, errors.Wrap(err, "parse config failed")
	}

	return &c, nil
}
