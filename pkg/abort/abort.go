package abort

import (
	"fmt"
	"os"
	"syscall"
)

func ByID(pid int) int {
	// Find process by PID.
	p, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Error: Process could not be found. ", p)
	}

	fmt.Println("Process INFO:", p)

	// Check if process actually exists.
	isProcessRunning(p)

	// Kill process
	p.Kill()

	// Check if process has been killed or it still runs.
	status := isProcessRunning(p)

	// Return status code & error if any
	if status != nil {
		return 1
	}
	return 0
}

func isProcessRunning(p *os.Process) error {
	err := p.Signal(syscall.Signal(0))
	if err != nil {
		fmt.Println("ERROR: Not found process:", err)
	}
	return err
}

// TODO Implement functionality to abort process by Name.
