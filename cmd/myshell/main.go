package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	execos "os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type cmdexec func(...string)

var commands = make(map[string]cmdexec)

func main() {
	loadCommands()

	reader := bufio.NewReader(os.Stdin)

	repl(reader)
}

func loadCommands() {
	commands["exit"] = func(s ...string) { exit(s) }
	commands["echo"] = func(s ...string) { echo(s) }
	commands["type"] = func(s ...string) { typeargs(s) }
	commands["cd"] = func(s ...string) { cd(s) }
	commands["pwd"] = func(_ ...string) { pwd() }
}

func repl(reader *bufio.Reader) {
	for {
		input := read(reader)

		cmd, args := parse(input)

		cmdfunc, ok := commands[cmd]

		if ok {
			cmdfunc(args...)
		} else {
			execFromPath(cmd, args)
		}
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

func parse(input string) (string, []string) {
	args := strings.Split(strings.TrimSpace(input), " ")
	return args[0], args[1:]
}

func exit(args []string) {
	if len(args) > 0 {
		exitCode, _ := strconv.Atoi(args[0])
		os.Exit(exitCode)
	} else {
		os.Exit(0)
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
	_, contains := commands[arg]

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

func execFromPath(cmd string, args []string) {
	execution := execos.Command(cmd, args...)

	output, err := execution.CombinedOutput()
	if err != nil {
		notFound(cmd)
	}

	fmt.Fprintf(os.Stdout, "%s", output)
}

func pwd() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", wd)
}

func cd(args []string) {
	var dir string
	if len(args) == 0 {
		dir = getHomeDir()
	} else {
		args[0] = strings.Replace(args[0], "~", getHomeDir(), -1)
		dir = args[0]
	}

	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", dir)
	}
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return home
}
