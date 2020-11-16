package main

import (
	"github.com/omnis-org/omnis-client/config"
	"github.com/omnis-org/omnis-client/internal/net"
	"github.com/omnis-org/omnis-client/internal/version"
	"github.com/omnis-org/omnis-client/pkg/client_informations"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	cfgCmdLine := struct {
		configFile string
	}{}

	cmdLine := kingpin.New(filepath.Base(os.Args[0]), "omnis-client")
	cmdLine.Version(version.BuildVersion)
	cmdLine.HelpFlag.Short('h')
	cmdLine.Flag("config.file", "Omnis configuration file path").Default("omnis.json").StringVar(&cfgCmdLine.configFile)

	_, err := cmdLine.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("Error parsing command line arguments : ", err)
	}

	config, err := config.LoadConfig(cfgCmdLine.configFile)
	if err != nil {
		log.Error("ParseConfig failed :", err)
	}

	for true {
		infos, err := client_informations.GetInformations(config.Client)
		if err != nil {
			log.Error("GetInformation failed :", err)
		}

		err = net.SendInformations(config.Server, infos)
		if err != nil {
			log.Error("SendInformations failed :", err)
		}

		d := time.Duration(config.Client.SendTime) * time.Second

		time.Sleep(d)
	}
}
