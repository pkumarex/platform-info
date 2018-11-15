package platforminfo

import "runtime"

// PlatformInfo is the interface that wraps the Windows or Linux
// implementation of platform information commands.
//
// These methods return error if there is error while running the command.
//
// BiosName() - Returns the BIOS OEM name (Currently not supported)
//
// BiosVersion() - Returns the BIOS version (Currently not supported)
//
// HardwareUUID() - Returns the Host's Hardware UUID
//
// OSName()  - Returns the Operating System(OS) name (Currently not supported)
//
// OSVersion() - Returns the Operating System(OS) version (Currently not supported)
//
// ProcessorFlags() - Returns the host CPU supported features/flags/instructions (Currently not supported)
//
// ProcessorInfo() - Returns the Processor(CPU) Information (Currently not supported)
//
// VMMName() - Returns the VMM(Hypervisor) name (Currently not supported)
//
// VMMVersion() - Returns the VMM(Hypervisor) version (Currently not supported)
//
// TPMVersion() - Returns the TPM Chip version (Currently not supported)
//
// HostName() - Returns Host's name (Currently not supported)
//
// NoOfSockets() - Returns the number of sockets (Currently not supported)
//
// TPMEnabled() - Returns true is TPM(Trusted Platform Module) of the platform is enabled (Currently not supported)
//
// TXTEnabled() - Returns true is TXT(Trusted Execution Technology) of the platform is enabled (Currently not supported)
//
type PlatformInfo interface {
	BiosName()         (string,error)
	BiosVersion()         (string,error)
	HardwareUUID()     (string,error)
	OSName()        (string,error)
	OSVersion()        (string,error)
	ProcessorFlags()   ([]string,error)
	ProcessorInfo()    (string,error)
	VMMName()          (string,error)
	VMMVersion()       (string,error)
	TPMVersion()       (string,error)
	HostName()         (string,error)
	NoOfSockets()      (int,error)
	TPMEnabled()       (bool,error)
	TXTEnabled()       (bool,error)
}

// This is the method that returns the implementation of platform info
// depending on the OS version.
// It returns PlatformInfo interface which then can be used to call a particular
// method to retrieve platform related information that are defined in PlatformInfo interface.
func New() PlatformInfo {
	if runtime.GOOS == "windows" {
		return new(windowsPlatformInfo)
	} else {
		return new(linuxPlatformInfo)
	}
}