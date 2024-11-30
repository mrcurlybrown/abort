package abort

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestAbortProcessByID(*testing.T) {
	// Start process to test & read output.
	cmd := exec.Command("/bin/bash", "", "./helpers/pid.sh")
	cmdErr := cmd.Start()
	if nil != cmdErr {
		fmt.Println("ERROR::Command failed", cmdErr)
	}
	p := cmd.Process
	if nil == p {
		fmt.Println("ERROR::Process failed", p)
	}

	// kill process by ID.
	ByID(p.Pid)

}

// TODO Write test to kill process by Name.
// Test to abort process by Name.
func TestAbortProcessByName() {
}
