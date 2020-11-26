package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/omnis-org/omnis-client/config"
	"github.com/omnis-org/omnis-client/pkg/client_informations"

	log "github.com/sirupsen/logrus"
)

func SendInformations(clientInfos *client_informations.Informations) error {
	timeout := time.Duration(config.GetConfig().Server.Timeout) * time.Second
	client := http.Client{Timeout: timeout}

	jsonClientInfos, err := json.Marshal(clientInfos)
	if err != nil {
		return fmt.Errorf("json.Marshal failed <- %v", err)
	}

	protocol := "http"
	if config.GetConfig().Server.TLS {
		protocol = "https"
	}

	url := fmt.Sprintf("%s://%s:%d/api/informations", protocol, config.GetConfig().Server.ServerIP, config.GetConfig().Server.ServerPort)

	log.Info("SendInformations : ", url, "\n", string(jsonClientInfos))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonClientInfos))
	if err != nil {
		return fmt.Errorf("http.NewRequest failed <- %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client.Do failed <- %v", err)
	}

	defer resp.Body.Close()

	return nil
}
