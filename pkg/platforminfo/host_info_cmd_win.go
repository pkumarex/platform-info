package platforminfo

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type HostInfoCmdWin struct {}

func (HostInfoCmdWin) retrieveBiosName() string {
	return ""
}

func (HostInfoCmdWin) retrieveHardwareUUID() string {
	cmd := exec.Command("wmic", "path", "Win32_ComputerSystemProduct", "get", "uuid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	result := strings.Split(string(out), "\n")
	hardwareUUID := ""
	if len(result) > 1 {
		re := regexp.MustCompile("\\s|\\r")
		hardwareUUID = re.ReplaceAllString(result[1], "")
	}
	return hardwareUUID
}

func (HostInfoCmdWin) retrieveOSVersion() string {
	return ""
}

func (HostInfoCmdWin) retrieveProcessorFlags() string {
	return ""
}

func (HostInfoCmdWin) retrieveProcessorInfo() string {
	return ""
}

func (HostInfoCmdWin) retrieveVMMName() string {
	return ""
}

func (HostInfoCmdWin) retrieveVMMVersion() string {
	return ""
}

func (HostInfoCmdWin) retrieveTpmVersion() string {
	return ""
}

func (HostInfoCmdWin) retrieveHostName() string {
	return ""
}

func (HostInfoCmdWin) retrieveNoOfSockets() string {
	return ""
}

func (HostInfoCmdWin) retrieveTPMEnabled() string {
	return ""
}

func (HostInfoCmdWin) retrieveTXTEnabled() string {
	return ""
}
