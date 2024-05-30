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

		eval(input)
	}
}

func read(reader *bufio.Reader) string {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimRight(input, "\n")
}

func eval(input string) {
	cmd, args := parse(input)

	exec(cmd, args)
}

func parse(input string) (string, []string) {
	args := strings.Split(strings.TrimSpace(input), " ")
	return args[0], args[1:]
}

func exec(cmd string, args []string) {
	switch cmd {
	case "exit":
		exit(args)
	case "echo":
		echo(args)
	default:
		notFound(cmd)
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

func echo(args []string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(args, " "))
}

func notFound(cmd string) {
	fmt.Printf("%s: command not found\n", cmd)
}
