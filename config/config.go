package config

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"sync"
)

/* This program contains the structures of data for the gathering of the machine information.
It also setup the configuration for the gathering of data
*/

type ServerConfig struct {
	Timeout            int64  `json:"timeout"`
	ServerIP           string `json:"server_ip"`
	ServerPort         int64  `json:"server_port"`
	TLS                bool   `json:"tls"`
	InsecureSkipVerify bool   `json:"insecure_skip_verify"`
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
	// TODO : Voir pour config en varible globale (bonne pratiques)
	var loadedConfigTmp Config
	fileContent, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile failed <- %v", err)
	}

	err = json.Unmarshal(fileContent, &loadedConfigTmp)
	if err != nil {
		return fmt.Errorf("json.Unmarshal failed <- %v", err)
	}
	loadedConfig = &loadedConfigTmp
	return nil
}

// TODO : A voir pour fichier json par default -- Default.json
func defaultConfig() *Config {
	sc := ServerConfig{10, "127.0.0.1", 4320, false, false}
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
