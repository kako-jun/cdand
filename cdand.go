package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/kako-jun/cdand/cdand-modules"
)

func parseArgs() (dirPath string, command string, commandArgs []string, err error) {
	flag.Parse()
	if flag.NArg() < 2 {
		err = errors.New("invalid argument")
		return
	}

	args := flag.Args()
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

	cdand.Exec(dirPath, command, args)
}
