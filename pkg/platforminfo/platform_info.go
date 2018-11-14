package platforminfo

import "runtime"

type platformInfo interface {
	BiosName() 			(string,error)
	HardwareUUID() 		(string,error)
	OSVersion() 		(string,error)
	ProcessorFlags() 	([]string,error)
	ProcessorInfo()		(string,error)
	VMMName() 			(string,error)
	VMMVersion() 		(string,error)
	TPMVersion() 		(string,error)
	HostName() 			(string,error)
	NoOfSockets() 		(int,error)
	TPMEnabled() 		(bool,error)
	TXTEnabled() 		(bool,error)
}

func New() platformInfo {
	if runtime.GOOS == "windows" {
		return new(windowsPlatformInfo)
	} else {
		return new(linuxPlatformInfo)
	}
}