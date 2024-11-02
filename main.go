package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Usage: %s run/update\n", args[0])
	}

	binary := args[0]
	mode := args[1]
	selfPid := os.Getpid()

	if mode == "run" {
		log.Println("Simulating work")
	} else if mode == "update" {
		// Using exec.Command + Run instead of syscall.Exec because syscall.Exec inherits PID and the other does not since it spawns a subprocess.
		execCmd := exec.Command("/Users/rase/dev/auto_updating/auto_update_updater/AutoUpdateUpdater", strconv.Itoa(selfPid), binary)
		if err := execCmd.Start(); err != nil {
			log.Fatal(err)
		}
		log.Println("Updating binary...")
		os.Exit(0)
	} else {
		log.Fatal("Unknown mode, expected run/update")
	}
}
