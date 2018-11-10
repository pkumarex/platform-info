package platforminfo

import (
	"fmt"
	"testing"
)

func TestHardwareUUID(t *testing.T) {
	p := NewPlatformInfo()
	fmt.Printf("Hardware UUID:%s\n", p.GetHardwareUUID())
}
