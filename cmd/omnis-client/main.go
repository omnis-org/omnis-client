package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/omnis-org/omnis-client/config"
	"github.com/omnis-org/omnis-client/internal/net"
	"github.com/omnis-org/omnis-client/internal/version"
	"github.com/omnis-org/omnis-client/pkg/client_informations"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func init() {
	cmdLine := kingpin.New(filepath.Base(os.Args[0]), "omnis-client")
	cmdLine.Version(version.BuildVersion)
	cmdLine.HelpFlag.Short('h')
	verbose := cmdLine.Flag("verbose", "Verbose mode.").Short('v').Bool()
	configFile := cmdLine.Arg("config.file", "Omnis configuration file path").Default("omnis.json").String()

	_, err := cmdLine.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("cmdLine.Parse failed <- ", err)
	}

	// logger
	log.SetFormatter(&nested.Formatter{
		HideKeys: true,
	})
	log.SetOutput(os.Stderr)
	if *verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	// config
	err = config.LoadConfig(configFile)
	if err != nil {
		log.Warn("config.LoadConfig failed <- ", err)
	}
}

func main() {
	// TODO : REFACTOR BOUCLE INFINI
	for true {
		infos, err := client_informations.GetInformations()
		if err != nil {
			log.Error("client_informations.GetInformation failed <- ", err)
		}

		//TODO : DON'T SEND INFO IF SAME STRUCT INFOS - NEW FUNCTION EQUALS FOR STRUCT
		err = net.SendInformations(infos)
		if err != nil {
			log.Error("net.SendInformations failed <- ", err)
		}

		time.Sleep(time.Duration(config.GetConfig().Client.SendTime) * time.Second)
	}
}
