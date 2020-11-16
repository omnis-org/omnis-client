package config

import (
	"io/ioutil"

	"github.com/tidwall/gjson"
)

type ServerConfig struct {
	Timeout    int64
	ServerIP   string
	ServerPort int64
	TLS        bool
}

type ClientConfig struct {
	Location  string
	Perimeter string
	SendTime  int64
}

type Config struct {
	Server *ServerConfig
	Client *ClientConfig
}

func LoadConfig(configFile string) (*Config, error) {
	json, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	// Load database config
	serverJSON := gjson.GetBytes(json, "server")
	serverConf := ServerConfig{serverJSON.Get("timeout").Int(),
		serverJSON.Get("server_ip").String(),
		serverJSON.Get("server_port").Int(),
		serverJSON.Get("tls").Bool()}
	clientJSON := gjson.GetBytes(json, "client")
	clientConf := ClientConfig{clientJSON.Get("location").String(),
		clientJSON.Get("perimeter").String(),
		clientJSON.Get("send_time").Int()}
	conf := Config{&serverConf, &clientConf}
	return &conf, nil
}
