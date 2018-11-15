package platforminfo

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

type linuxPlatformInfo struct {}

func (linuxPlatformInfo) BiosName() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) BiosVersion() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) HardwareUUID() (string,error) {
	//Output of command:
	//4219B2F5-C25F-6AF2-573C-35B0DF557236
	cmd := exec.Command("dmidecode", "-s", "system-uuid")
	out, err := cmd.CombinedOutput()
	hardwareUUID := ""
	if err != nil {
		//print error and exit
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// Split the output separated by new line into list
	result := strings.Split(string(out), "\n")
	for i := range result {
		// If first few lines that start with prefix #, ignore as # indicates comments
		if strings.HasPrefix(strings.TrimSpace(result[i]), "#") {
			continue
		}
		// remove spaces
		hardwareUUID = strings.TrimSpace(result[i])
		break
	}
	return hardwareUUID,nil
}

func (linuxPlatformInfo) OSName() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) OSVersion() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) ProcessorFlags() ([]string,error) {
	return []string{},errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) ProcessorInfo() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) VMMName() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) VMMVersion() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) TPMVersion() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) HostName() (string,error) {
	return "",errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) NoOfSockets() (int,error) {
	return -1,errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) TPMEnabled() (bool,error) {
	return false,errors.New("unsupported command and will be supported in future")
}

func (linuxPlatformInfo) TXTEnabled() (bool,error) {
	return false,errors.New("unsupported command and will be supported in future")
}

