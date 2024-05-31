[![progress-banner](https://backend.codecrafters.io/progress/shell/a0af3094-777c-45ad-92b6-431fccefdeac)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# Build your own Shell

This is my solution in GO for the CodeCrafter's
["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview).

### Overview

This challenge consist in building a simple repl shell in any desired language featuring the commands

```sh
exit code
echo something
type any_command
cd {~, ./some_relative_path, /some/absolute/path}
./executable.sh
```

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

### Running my shell on your machine

To run the shell you must have go version `1.22.3` installed on your machine. If by any chance you use asdf the `.tool-versions` is already configured

Once installed, simply run the following command:

```sh
./your_shell.sh
```

It will start the custom shell and you can start typing commands.

**Note**: any command from the user `$PATH` is valid, but it will attempt to run builtin first
