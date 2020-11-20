package config

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"sync"
)

type ServerConfig struct {
	Timeout    int64  `json:"timeout"`
	ServerIP   string `json:"server_ip"`
	ServerPort int64  `json:"server_port"`
	TLS        bool   `json:"tls"`
}

type ClientConfig struct {
	Location  string `json:"location"`
	Perimeter string `json:"perimeter"`
	SendTime  int64  `json:"send_time"`
}

type Config struct {
	Server *ServerConfig `json:"server"`
	Client *ClientConfig `json:"client"`
}

var lockConfig = &sync.Mutex{}
var loadedConfig *Config = nil

func LoadConfig(configFile *string) error {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	var loadedConfigTmp Config
	jsonS, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile failed <- %v", err)
	}

	err = json.Unmarshal(jsonS, &loadedConfigTmp)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	loadedConfig = &loadedConfigTmp
	return nil
}

func defaultConfig() *Config {
	sc := ServerConfig{10, "127.0.0.1", 4320, false}
	cc := ClientConfig{"default_network", "default_perimeter", 60}
	return &Config{&sc, &cc}
}

func GetConfig() *Config {
	lockConfig.Lock()
	defer lockConfig.Unlock()
	if loadedConfig == nil {
		return defaultConfig()
	}
	return loadedConfig
}
