package core

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetMachineHostName() (string, error) {
	return os.Hostname()
}

func GetKernelInformation() ([]string, error) {
	cmd := exec.Command("uname", "-srm")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Warn("getInfo:", err)
	}
	osStr := strings.Replace(out.String(), "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")
	osInfo = append(osInfo, runtime.GOOS)
	osInfo = append(osInfo, string(runtime.NumCPU()))
	log.Info("OS: ", osInfo[0])
	log.Info("Core: ", osInfo[1])
	log.Info("Platform: ", osInfo[2])
	log.Info("GoOs: ", osInfo[3])
	log.Info("CPUs: ", osInfo[4])
	return osInfo, nil
}
