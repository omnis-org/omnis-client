package core

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetMachineHostName() (string, error) {
	return os.Hostname()
}

func GetKernelInformation() (SystemInformation, error) {
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
	hostname, err := GetMachineHostName()
	if err != nil {
		return SystemInformation{}, err
	}
	sysInfo := SystemInformation{
		OS:          osInfo[0],
		HostName:    hostname,
		Platform:    osInfo[2],
		Core:        osInfo[1],
		GoOsVersion: runtime.GOOS,
		CPU:         strconv.Itoa(runtime.NumCPU()),
	}
	return sysInfo, nil
}
