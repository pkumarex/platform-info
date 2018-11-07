package cmd

import (
	"runtime"
)

type ICommandFactory struct {
}

func (f ICommandFactory) getCmd() ICommand {
	var iCmd ICommand
	if runtime.GOOS == "windows" {
		iCmd = new(HostInfoCmdWin)
	} else {
		iCmd = new(HostInfoCmd)
	}
	return iCmd
}