package config

import (
	"encoding/json"
	"os"
	"time"
)

type serverConfig struct {
	Addr                   string        `json:"addr"`
	ReadTimeout            time.Duration `json:"readTimeout"`
	WriteTimeout           time.Duration `json:"writeTimout"`
	IdleTimeout            time.Duration `json:"idleTimeout"`
	MultipartFormMaxMemory int64         `json:"maxMemory"`
}

var Server serverConfig

func Load(filename string) error {

	cfgFile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(cfgFile, &Server)

	if err != nil {
		return err
	}

	return nil
}
