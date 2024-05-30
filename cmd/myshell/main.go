package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	repl(reader)

	os.Exit(0)
}

func repl(reader *bufio.Reader) {
	for {
		input := read(reader)

		cmd := eval(input)

		if valid := validateCommand(cmd); !valid {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}

func read(reader *bufio.Reader) string {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading from buffer %s", err)
		os.Exit(-1)
	}

	return input
}

func eval(input string) string {
	return strings.TrimRight(input, "\n") // TODO: add proper evaluation
}

func validateCommand(cmd string) bool {
	return false // TODO: add proper validation
}
