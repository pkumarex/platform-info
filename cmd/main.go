package main

import (
	"fmt"
	"lib-go-platform-info/pkg/platforminfo"
)

func main() {
	p := platforminfo.NewPlatformInfo()
	fmt.Printf("Hardware UUID:%s\n", p.GetHardwareUUID())
}
