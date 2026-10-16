package freeport

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// freePort finds and terminates any existing process running on the specified port
// so the server can always start on the same port without "address already in use" errors.
func FreePort(port string) {
	// Remove leading colon if present (e.g. ":8080" becomes "8080")
	portNum := strings.TrimPrefix(port, ":")

	// Prepare shell command "lsof -t -i:8080" to look up process IDs (PIDs) using this port
	// -t: terse output (prints only process IDs)
	// -i: selects network socket matching the port
	cmd := exec.Command("lsof", "-t", "-i:"+portNum)

	// Run command and capture output (contains PID list as text)
	output, err := cmd.Output()
	if err != nil {
		// If command returns error or no process is found, return safely
		return
	}

	// Get current process ID so we don't accidentally terminate our own program
	currentPID := os.Getpid()

	// Split the output string into individual PID strings (handles newlines/spaces)
	pids := strings.Fields(string(output))

	// Loop through each process ID found on this port
	for _, pidStr := range pids {
		// Convert the PID from string to integer
		pid, err := strconv.Atoi(pidStr)

		// Make sure conversion succeeded and that the PID is not our own process
		if err == nil && pid != currentPID {
			// Locate the OS process by its PID
			if proc, err := os.FindProcess(pid); err == nil {
				// Print information about the process being terminated
				fmt.Printf("Freeing port %s (killing old process PID %d)...\n", portNum, pid)

				// Terminate the process (sends SIGKILL)
				_ = proc.Kill()

				// Wait briefly (200ms) to allow the OS network stack to release the socket
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}