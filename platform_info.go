// +build darwin linux

package platforminfo

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

// BiosName retrieves the host BIOS name
// Unsupported
func BiosName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// BiosVersion retrieves the host BIOS version
// Unsupported
func BiosVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// HardwareUUID retireves the host hardware UUID.
// An example UUID is 4219B2F5-C25F-6AF2-573C-35B0DF557236
func HardwareUUID() (string, error) {
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
	return hardwareUUID, nil
}

// OSName retrieves the host OS name
// Unsupported
func OSName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// OSVersion retrieves the host OS version
// Unsupported
func OSVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// ProcessorFlags retrieves the processor flags (CPUID)
// Unsupported
func ProcessorFlags() ([]string, error) {
	return []string{}, errors.New("unsupported command and will be supported in future")
}

// ProcessorInfo retrieves the host processor info
// Unsupported
func ProcessorInfo() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// VMMName retrives the name of the hypervisor running on the host
// Unsupported
func VMMName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// VMMVersion retrieves the version of the hypervisor running on the host
// Unsupported
func VMMVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// TPMVersion retrieves the version of the installed Trusted Platform Module (TPM)
// Unsupported
func TPMVersion() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// HostName retireves the network hostname
// Unsupported
func HostName() (string, error) {
	return "", errors.New("unsupported command and will be supported in future")
}

// NoOfSockets retrieves the number of CPU sockets on the host platform
func NoOfSockets() (int, error) {
	return -1, errors.New("unsupported command and will be supported in future")
}

// TPMEnabled indicates whether the Trusted Platform Module is enabled or not
// Unsupported
func TPMEnabled() (bool, error) {
	return false, errors.New("unsupported command and will be supported in future")
}

// TXTEnabled indicates whether Intel TXT is enabled or not
// Unsupported
func TXTEnabled() (bool, error) {
	return false, errors.New("unsupported command and will be supported in future")
}
