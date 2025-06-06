# gun Usage

## Without Parameters
Gun compiles the program and runs it, using `build` and `run` defined in the local `./commands`
file, and if not found, then defaults defined in `config.toml`.
```console
$ gun [run args]

FLAGS:
  --flags [Flags]    Specify optional build flags
  --global           Run the command using the global config, overriding the local config
```

## Init
Initiate a new project with templates, licensing, and a structured layout. The data flow of the command proceeds as follows:
- `config.toml` is read and all configs are loaded.
- If the `--here` flag is not specified, a new directory is created, and gun cd into it.
- Unless `copy_config` is set to false, all defined [xxx.dev] commands are copied from the local `./commands` file.
- All specified `directories` are created.
- All designated `files` are created.
- All defined `templates` are copied over.
- The specified `license` file is copied.
- The designated `command` is executed.
- A Git repository is initialised.
- The `.gitignore` file is generated.
- A Git commit is made with the specified `git_message`
```console
$ gun init <project name> <project language> [flags]

FLAGS:
  --here             Create the project in the current dir, instead of making a new one
  --no-git           Create the project but not git repository
  --no-templates     Do not copy any config defined templates to the project
  --license [title]  Use the specified license for the project instead of the config defined

EXAMPLES:
  # Create a new C project titled FooBar
  $ gun init FooBar c

  # Create a new python project titled HelloWorld in current working directory
  $ gun init HelloWorld py --here

  # Create a new ESP32 project titled BlinkLED without any defined templates with GPLv2 license
  $ gun init BlinkLED esp --license GPLv2 --no-templates
```

## Template
Manage project templates (add, create, list)
```console
$ gun template [arguments] [flags]

FLAGS:
  --list             List all saved template
  --print-dir        Directory where all files are saved
  --new [title]      Create a new template
  --add [title]      Add a saved template to the current project
  --manage           Manage all saved templates
```

## License
Manage license files (create, replace, list)
```console
$ gun license [arguments] [flags]

FLAGS:
  --list             List all saved license
  --print-dir        Directory where all files are saved
  --new [title]      Create a new license
  --replace [title]  Replace the current license with a new one

```

## Run
Run the program with optional arguments.
```console
$ gun run [arguments]

EXAMPLES:
  # Run the last compiled Go program
  $ gun run

  # Run the last compiled C program
  $ gun run

  # Run the last compiled program with `"hello world" 1234` passed as arguments
  $ gun run "hello world" 1234
```

## Build
Compile the program with optional flags.
```console
$ gun build [arguments] [flags]

FLAGS:
  --global           Run the command using the global config, overriding the local config

EXAMPLES:
  # Build the Go project
  $ gun build

  # Build the Java project with `-g:none` parameters
  $ gun build -g:none

  # Build the C flag with default configs and with -Wextra -std=c11 parameters
  $ gun build -Wextra -std=c11 --global
```

## Clean
The command to clean project build files.
```console
$ gun clean [arguments] [flags]

FLAGS:
  --global           Run the command using the global config, overriding the local config

EXAMPLES:
  # Clean the project
  $ gun clean

  # Clean the project as defined in global config
  $ gun clean --global

```

## Debug
Run the configured debugger for the project.
```console
$ gun debug [arguments] [flags]

FLAGS:
  --global           Run the command using the global config, overriding the local config

EXAMPLES:
  # Run the set debugger for the project
  $ gun debug

  # Run the default debugger with `-batch -ex "run" -ex "bt"` arguments
  $ gun debug -batch -ex "run" -ex "bt" --global
```

## Test
Run the test suite.
```console
$ gun test [arguments] [flags]

FLAGS:
  --global           Run the command using the global config, overriding the local config

EXAMPLES:
  # Run the defined test suite
  $ gun test

  # Run the global defined test suit with `"foo" "bar"` arguments
  $ gun test foo bar --global
```

## Config
Edit or create a local project configuration file.
```console
$ gun config [arguments] [flags]

FLAGS:
  --local            Edit the local config file instead of the global one

EXAMPLES:
  # Update global config
  $ gun config

  # Update local config
  $ gun config --local
```

## Help
Get information about any and all commands.
```console
 $ gun help [command]

EXAMPLES:
  # Get general information of the program
  $ gun help

  # Get all supported flags and usage of "init" subcommand
  $ gun help init
```

## Version
Get current gun version.
```console
$ gun version
```

