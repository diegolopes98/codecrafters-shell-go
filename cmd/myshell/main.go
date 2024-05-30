package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading from buffer %s", err)
		os.Exit(-1)
	}

	cmd := strings.TrimRight(input, "\n")
	if valid := validateCommand(cmd); !valid {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
	}

	os.Exit(0)
}

func validateCommand(cmd string) bool {
	return false // TODO: add proper validation
}
