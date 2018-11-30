package platforminfo

import (
	"log"
	"testing"
)

func TestHardwareUUID(t *testing.T) {
	p := New()
	hardwareUUID,err := p.HardwareUUID()
	if err != nil {
		log.Println(err))
	} else {
		log.Println("Hardware UUID:\n", hardwareUUID)
	}
}
