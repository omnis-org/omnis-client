package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/omnis-org/omnis-client/config"
	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"time"

	log "github.com/sirupsen/logrus"
)

func SendInformations(serverConfig *config.ServerConfig, infos *client_informations.Informations) error {
	time := time.Duration(serverConfig.Timeout) * time.Second
	client := http.Client{Timeout: time}

	jsonInfos, err := json.Marshal(infos)
	if err != nil {
		return err
	}

	protocol := "http"
	if serverConfig.TLS {
		protocol = "https"
	}

	url := fmt.Sprintf("%s://%s:%d/api/informations", protocol, serverConfig.ServerIP, serverConfig.ServerPort)

	log.Info("SendInformations : ", url, "\n", string(jsonInfos))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonInfos))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
