package config

import (
	"encoding/json"
	"fmt"
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

	fmt.Println(os.Getwd())

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
