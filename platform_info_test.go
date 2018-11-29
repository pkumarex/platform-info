package platforminfo

import (
	"fmt"
	"testing"
)

func TestHardwareUUID(t *testing.T) {
	p := New()
	hardwareUUID,err := p.HardwareUUID()
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Hardware UUID:%s\n", hardwareUUID)
	}
}
