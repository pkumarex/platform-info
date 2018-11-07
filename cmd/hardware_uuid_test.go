package cmd

import (
	"fmt"
	"testing"
)

func TestStuff(t *testing.T) {
	p := NewPlatformInfo()
	fmt.Printf("Hardware UUID:%s\n", p.GetHardwareUUID())
}
