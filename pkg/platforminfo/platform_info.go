package platforminfo

type platformInfo struct {
	cmd ICommand
}

func NewPlatformInfo() *platformInfo {
	p := new(platformInfo)
	p.cmd = new(ICommandFactory).getCmd()
	return p
}

func GetBiosName() string {
	return ""
}

func (p platformInfo) GetHardwareUUID() string {
	return p.cmd.retrieveHardwareUUID()
}

func GetOSVersion() string {
	return ""
}

func GetProcessorFlags() string {
	return ""
}

func GetProcessorInfo() string {
	return ""
}

func GetVMMName() string {
	return ""
}

func GetVMMVersion() string {
	return ""
}

func GetTPMVersion() string {
	return ""
}

func GetHostName() string {
	return ""
}

func GetNoOfSockets() string {
	return ""
}
