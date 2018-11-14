package platforminfo

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

type linuxPlatformInfo struct {}

func (linuxPlatformInfo) BiosName() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) HardwareUUID() (string,error) {
	cmd := exec.Command("dmidecode", "-s", "system-uuid")
	out, err := cmd.CombinedOutput()
	hardwareUUID := ""
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return "",err
	}
	result := strings.Split(string(out), "\n")
	for i := range result {
		if strings.HasPrefix(strings.TrimSpace(result[i]), "#") {
			continue
		}
		hardwareUUID = strings.TrimSpace(result[i])
		break
	}
	return hardwareUUID,nil
}

func (linuxPlatformInfo) OSVersion() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) ProcessorFlags() ([]string,error) {
	return []string{},errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) ProcessorInfo() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) VMMName() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) VMMVersion() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) TPMVersion() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) HostName() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) NoOfSockets() (int,error) {
	return -1,errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) TPMEnabled() (bool,error) {
	return false,errors.New("unsupported command. will be supported in future")
}

func (linuxPlatformInfo) TXTEnabled() (bool,error) {
	return false,errors.New("unsupported command. will be supported in future")
}

