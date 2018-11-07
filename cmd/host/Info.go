package host

type Info struct {
	HostName string
	BiosName string
	BiosVersion string
	HardwareUUID string
	OSName string
	OSVersion string
	ProcessorInfo string
	VMMName string
	VMMVersion string
	ProcessorFlags string
	TPMVersion string
	PCRBanks []string
	NoOfSockets string
	TPMEnabled string
	TXTEnabled string
}
