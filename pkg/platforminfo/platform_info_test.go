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

//func fakeExecCommand(command string, args...string) *exec.Cmd {
//	cs := []string{"-test.run=TestHelperProcess", "--", command}
//	cs = append(cs, args...)
//	cmd := exec.Command(os.Args[0], cs...)
//	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
//	return cmd
//}
//
////const hardwareUUIDResult = "C8C8411F-F0CB-11E5-8343-9025330C6062"
//const mocked = "C8C8411F-F0CB-11E5-8343-9025330C7062"
//const exp = "C8C8411F-F0CB-11E5-8343-9025330C7062"
//func TestRunHardwareUUID(t *testing.T) {
//	p := New()
//	ExecCommand = fakeExecCommand
//	defer func(){ ExecCommand = exec.Command }()
//	out, err := p.HardwareUUID()
//	if err != nil {
//		t.Errorf("Expected nil error, got %#v", err)
//	}
//	if string(out) != exp {
//		t.Errorf("Expected %q, got %q", exp, out)
//	}
//}
//
//func TestHelperProcess(t *testing.T){
//	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
//		return
//	}
//	// some code here to check arguments perhaps?
//	fmt.Fprintf(os.Stdout, mocked)
//	os.Exit(0)
//}
