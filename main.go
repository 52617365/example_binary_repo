package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
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
	log.Printf("PID: %d\n", selfPid)

	if mode == "run" {
		log.Println("Simulating work")
	} else if mode == "update" {
		// Using exec.Command + Start instead of syscall.Exec because syscall.Exec inherits PID and the other does not since it spawns a subprocess.
		var pathToUpdater string
		if runtime.GOOS == "windows" {
			pathToUpdater = "C:\\Users\\rase\\Desktop\\AutoUpdateUpdater.exe"
		} else {
			pathToUpdater = "/Users/rase/dev/auto_updating/auto_update_updater/AutoUpdateUpdater"
		}
		execCmd := exec.Command(pathToUpdater, strconv.Itoa(selfPid), binary)
		if err := execCmd.Start(); err != nil {
			log.Fatal(err)
		}
		log.Println("Updating binary...")
		os.Exit(0)
	} else {
		log.Fatal("Unknown mode, expected run/update")
	}
}
