package cmd

import (
	"lib-go-platform-info/cmd/host"
)

type PlatformInfo struct {
	hostInfo host.Info
	cmd ICommand
}

func NewPlatformInfo() *PlatformInfo {
	p := new(PlatformInfo)
	p.cmd = new(ICommandFactory).getCmd()
	return p
}

func GetBiosName() string {
	return ""
}

func (p PlatformInfo) GetHardwareUUID() string {
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
