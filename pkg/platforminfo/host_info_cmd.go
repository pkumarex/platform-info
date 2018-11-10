package platforminfo

import (
	"log"
	"os/exec"
	"strings"
)

type HostInfoCmd struct {}

func (HostInfoCmd) retrieveBiosName() string {
	return ""
}

func (HostInfoCmd) retrieveHardwareUUID() string {
	cmd := exec.Command("dmidecode", "-s", "system-uuid")
	out, err := cmd.CombinedOutput()
	hardwareUUID := ""
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	result := strings.Split(string(out), "\n")
	for i := range result {
		if strings.HasPrefix(strings.TrimSpace(result[i]), "#") {
			continue
		}
		hardwareUUID = strings.TrimSpace(result[i])
		break
	}
	return hardwareUUID
}

func (HostInfoCmd) retrieveOSVersion() string {
	return ""
}

func (HostInfoCmd) retrieveProcessorFlags() string {
	return ""
}

func (HostInfoCmd) retrieveProcessorInfo() string {
	return ""
}

func (HostInfoCmd) retrieveVMMName() string {
	return ""
}

func (HostInfoCmd) retrieveVMMVersion() string {
	return ""
}

func (HostInfoCmd) retrieveTpmVersion() string {
	return ""
}

func (HostInfoCmd) retrieveHostName() string {
	return ""
}

func (HostInfoCmd) retrieveNoOfSockets() string {
	return ""
}

func (HostInfoCmd) retrieveTPMEnabled() string {
	return ""
}

func (HostInfoCmd) retrieveTXTEnabled() string {
	return ""
}

