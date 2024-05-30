package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	repl(reader)
}

func repl(reader *bufio.Reader) {
	for {
		input := read(reader)

		cmd, args := eval(input)

		valid := validateCommand(cmd)

		if cmd == "exit" {
			exit(args)
		}

		if !valid {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}

func read(reader *bufio.Reader) string {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return input
}

func eval(input string) (string, []string) {
	args := strings.Split(strings.TrimRight(input, "\n"), " ")
	return args[0], args[1:]
}

func validateCommand(cmd string) bool {
	switch cmd {
	case "exit":
		return true
	default:
		return false
	}
}

func exit(args []string) {
	if len(args) > 0 {
		exitCode, _ := strconv.Atoi(args[0])
		os.Exit(exitCode)
	} else {
		os.Exit(0) // TODO: maybe -1 better?
	}
}
