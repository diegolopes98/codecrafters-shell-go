package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	execos "os/exec"
	"path/filepath"
	"slices"
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
	case "type":
		typeargs(args)
	default:
		fullPath, isFromPath := isFromPath(cmd)
		if isFromPath {
			execFromPath(fullPath, args)
		} else {
			notFound(cmd)
		}
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

func typeargs(args []string) {
	for _, arg := range args {
		typearg(arg)
	}
}

func typearg(arg string) {
	builtin := []string{"exit", "echo", "type"} // TODO: review replication of commands check

	contains := slices.Contains(builtin, arg)
	if contains {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", arg)
	} else {
		fullPath, fromPath := isFromPath(arg)
		if fromPath {
			fmt.Fprintf(os.Stdout, "%s is %s\n", arg, fullPath)
		} else {
			fmt.Fprintf(os.Stdout, "%s not found\n", arg)
		}
	}
}

func isFromPath(arg string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		fullPath := filepath.Join(path, arg)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, true
		}
	}
	return "", false
}

func execFromPath(binPath string, args []string) {
	execution := execos.Command(binPath, args...)

	output, err := execution.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%s", output)
}
