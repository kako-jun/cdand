package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kako-jun/cdand/cdand-core"
)

// app version
var Version string = "1.0.0"

func parseArgs() (dirPath string, command string, commandArgs []string, err error) {
	var (
		versionFlag bool
	)

	flag.BoolVar(&versionFlag, "version", false, "print version number")

	flag.Parse()
	args := flag.Args()

	if versionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}

	if flag.NArg() < 2 {
		err = errors.New("invalid argument")
		return
	}

	dirPath = args[0]
	command = args[1]
	commandArgs = args[2:]
	return
}

// entry point
func main() {
	dirPath, command, args, err := parseArgs()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("usage:")
		fmt.Println("  cdand [dir] [command]")
		fmt.Println("  cdand [dir] [command] [command options]")
		return
	}

	cdand.Exec(dirPath, command, args...)
}
