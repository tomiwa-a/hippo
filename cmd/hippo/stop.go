package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running Hippo engine",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("hippo.pid")
		if err != nil {
			color.Red("Error: Hippo is not running (no hippo.pid found)")
			return
		}

		pid, err := strconv.Atoi(string(data))
		if err != nil {
			color.Red("Error: Invalid PID file")
			return
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			color.Red("Error: Could not find process %d", pid)
			return
		}

		err = process.Signal(syscall.SIGTERM)
		if err != nil {
			color.Red("Error: Failed to send stop signal: %v", err)
			return
		}

		fmt.Printf("Stopping Hippo (PID: %d)...\n", pid)
		os.Remove("hippo.pid")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
