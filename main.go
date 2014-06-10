package main

import (
	"errors"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [option]. Where [option] could be one of boot/build/versions/shell\n", os.Args[0])
	os.Exit(2)
}

func main() {
	buildDirectory, _ := os.Getwd()
	if len(os.Args) == 1 {
		usage()
	}
	for _, subCommand := range os.Args {
		if subCommand == "grit" {
			continue
		}
		if subCommand == "boot" {
			fmt.Println("Bootstrapping ", buildDirectory)
			return
		} else if subCommand == "build" {
			fmt.Println("Building ", buildDirectory)
			return
		} else if subCommand == "versions" {
			fmt.Println("Printing the versions of ", buildDirectory)
			return
		} else if subCommand == "shell" {
			fmt.Println("Opening shell for ", buildDirectory)
			return
		} else {
			fmt.Print(errors.New("Usage: grit [option]. Where [option] could be one of boot/build/versions/shell\n"))
			return
		}
	}
}
