package cmd

import (
	"time"

	"omnis-client/core"

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
	infosCmd.Flags().BoolP("network", "k", false, "gather network infos")
	rootCmd.AddCommand(infosCmd)
}

func gatherInfos(cmd *cobra.Command, args []string) error {
	begin := time.Now()
	// TODO Create interface to make object gatherer
	// Create in core new gatherInfos function
	gatherer := core.NewGatherer()
	result, err := gatherer.GatherMachineInfos(cmd.Context())
	if err != nil {
		return err
	}
	log.Info(result)
	log.Info("Scan execution time: %s", time.Since(begin))
	return nil
}
