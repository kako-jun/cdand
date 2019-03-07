// Package cdand is ***
package cdand

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// CdAnd is ***
type CdAnd struct {
	// instance variables
}

func (cdand CdAnd) exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func (cdand CdAnd) execCommand(command string, args ...string) (err error) {
	// fmt.Println(command)
	// fmt.Println(args)

	cmd := exec.Command(command, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

	outScanner := bufio.NewScanner(stdout)
	for outScanner.Scan() {
		fmt.Fprintln(os.Stdout, outScanner.Text())
	}

	errScanner := bufio.NewScanner(stderr)
	for errScanner.Scan() {
		fmt.Fprintln(os.Stderr, errScanner.Text())
	}

	cmd.Wait()
	return
}

func (cdand CdAnd) start(dirPath string, command string, args ...string) (err error) {
	if !cdand.exists(dirPath) {
		err = errors.New(dirPath + " not found.")
		return
	}

	currentDirPath, _ := os.Getwd()
	os.Chdir(dirPath)

	err = cdand.execCommand(command, args...)

	os.Chdir(currentDirPath)
	return
}

// Exec is ***
func Exec(dirPath string, command string, args ...string) (errReturn error) {
	cdand := new(CdAnd)
	if err := cdand.start(dirPath, command, args...); err != nil {
		fmt.Println("error:", err)
		errReturn = errors.New("error")
		return
	}

	return
}
