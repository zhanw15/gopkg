// Package easyssh provides a simple implementation of some SSH protocol
// features in Go. You can simply run a command on a remote server or get a file
// even simpler than native console SSH client. You don't need to think about
// Dials, sessions, defers, or public keys... Let easyssh think about it!
package easyssh

import (
	"bufio"
	"io"
	"os/exec"
)


// Stream returns one channel that combines the stdout and stderr of the command
// as it is run on the remote machine, and another that sends true when the
// command is done. The sessions and channels will then be closed.
func Stream(command string) (output chan string, done chan bool, err error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	// connect to both outputs (they are of type io.Reader)
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return output, done, err
	}
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return output, done, err
	}
	// combine outputs, create a line-by-line scanner
	outputReader := io.MultiReader(outReader, errReader)
	err = cmd.Start()
	scanner := bufio.NewScanner(outputReader)
	// continuously send the command's output over the channel
	outputChan := make(chan string)
	done = make(chan bool)
	go func(scanner *bufio.Scanner, out chan string, done chan bool) {
		defer close(outputChan)
		defer close(done)
		for scanner.Scan() {
			outputChan <- scanner.Text()
		}
		// close all of our open resources
		done <- true
	}(scanner, outputChan, done)
	return outputChan, done, err
}

// Runs command on remote machine and returns its stdout as a string
func Run(command string) (outStr string, err error) {
	outChan, doneChan, err := Stream(command)
	if err != nil {
		return outStr, err
	}
	// read from the output channel until the done signal is passed
	stillGoing := true
	for stillGoing {
		select {
		case <-doneChan:
			stillGoing = false
		case line := <-outChan:
			outStr += line + "\n"
		}
	}
	// return the concatenation of all signals from the output channel
	return outStr, err
}
