package platforminfo

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type windowsPlatformInfo struct {}

func (windowsPlatformInfo) BiosName() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) HardwareUUID() (string,error) {
	cmd := exec.Command("wmic", "path", "Win32_ComputerSystemProduct", "get", "uuid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return "",err
	}
	result := strings.Split(string(out), "\n")
	hardwareUUID := ""
	if len(result) > 1 {
		re := regexp.MustCompile("\\s|\\r")
		hardwareUUID = re.ReplaceAllString(result[1], "")
	}
	return hardwareUUID,nil
}

func (windowsPlatformInfo) OSVersion() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) ProcessorFlags() ([]string,error) {
	return []string{},errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) ProcessorInfo() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) VMMName() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) VMMVersion() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) TPMVersion() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) HostName() (string,error) {
	return "",errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) NoOfSockets() (int,error) {
	return -1,errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) TPMEnabled() (bool,error) {
	return false,errors.New("unsupported command. will be supported in future")
}

func (windowsPlatformInfo) TXTEnabled() (bool,error) {
	return false,errors.New("unsupported command. will be supported in future")
}
