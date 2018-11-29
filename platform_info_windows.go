// +build windows

package platforminfo

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func BiosName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func BiosVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func hardwareUUID() (string, error) {
	//Output of command:
	//UUID
	//C8C8411F-F0CB-11E5-8343-9025330C6062
	//
	cmd := exec.Command("wmic", "path", "Win32_ComputerSystemProduct", "get", "uuid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		//print error and exit
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// Split the output separated by new line into list
	result := strings.Split(string(out), "\n")
	hardwareUUID := ""
	if len(result) > 1 {
		// remove all spaces from the second line as that line consists hardware uuid
		re := regexp.MustCompile("\\s|\\r")
		hardwareUUID = re.ReplaceAllString(result[1], "")
	}
	return hardwareUUID, nil
}

func OSName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func OSVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func ProcessorFlags() ([]string, error) {
	return []string{}, errors.New("unsupported command and will be supported in future")
}

func ProcessorInfo() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func VMMName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func VMMVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func TPMVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func HostName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

func NoOfSockets() (int, error) {
	return -1, errors.New("unsupported command and will be supported in future")
}

func TPMEnabled() (bool, error) {
	return false, errors.New("unsupported command and will be supported in future")
}

func TXTEnabled() (bool, error) {
	return false, errors.New("unsupported command and will be supported in future")
}
