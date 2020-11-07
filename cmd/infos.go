package cmd

import (
	"fmt"
	"time"

	"omnis-client/core"
	"omnis-client/internal/formatting"
	"omnis-client/internal/httppost"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// infosCmd represents the infos command

func init() {
	var infosCmd = &cobra.Command{
		Use:   "infos",
		Short: "Gather informations of system and network of the machine",
		Long:  `Gather informations of system and network of the machine`,
		RunE:  gatherInfos,
	}
	infosCmd.Flags().StringP("serverIP", "s", "", "Main Server IP")
	infosCmd.Flags().IntP("timeout", "t", 10, "Timeout for the HTTP requests (default: 10s)")
	rootCmd.AddCommand(infosCmd)
}

func gatherInfos(cmd *cobra.Command, args []string) error {
	config, err := parseConfig(cmd, args)
	begin := time.Now()
	fetcher := httppost.NewFetcher(config.HTTP.Timeout)
	gatherer := core.NewGatherer(fetcher)
	result, err := gatherer.GatherMachineInfos(cmd.Context())
	if err != nil {
		return err
	}
	log.Info(result)
	log.Info(fmt.Sprintf("Scan execution time: %s", time.Since(begin)))
	formatting.ExportJSON(result)
	if config.ServerIP != "" {
		resp, err := gatherer.SendInfos(config.ServerIP, result)
		if err != nil {
			return err
		}
		log.Info(resp)
	}
	return nil
}

func parseConfig(cmd *cobra.Command, args []string) (*core.Config, error) {
	serverIP, err := cmd.Flags().GetString("serverIP")
	if err != nil {
		return nil, fmt.Errorf("invalid value for serverIP: %v", err)
	}
	timeout, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return nil, fmt.Errorf("Invalid value for timeout: %v", err)
	}
	config := &core.Config{
		ServerIP: serverIP,
		HTTP: core.HTTPConfig{
			Timeout: timeout,
		},
	}
	return config, nil
}
