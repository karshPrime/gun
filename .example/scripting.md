
# Scripting Guide

The purpose of scripting support is to make the `[init]` configurations more dynamic within the
project. It is important to note that this syntax is only valid for `[init]` and not for `[dev],` as
much of the project information is not parsed when executing `[dev]` commands.


## `scripts/` Directory

Users can save shell (or batch) scripts in the `<GUN-CONFIG>/scripts` directory, such as
`~/.config/gun/scripts/.` In the `config.toml` file, these scripts can be referenced using the
`%SCRIPT%` placeholder.

For example, the `cmakegenerate.sh` file takes two arguments: the project name and the programming
language (C/C++). It uses these arguments to generate the appropriate CMake configuration for the
project, running from the project's root directory. The file is saved as
`~/.config/gun/scripts/cmakegenerate.sh`, and in the config, it is called as follows:
```toml
[init.c]
command = "%SCRIPT% cmakegenerate.sh %PROJECT_NAME% c"
```


## Project Name

To use with `%SCRIPT%` or within the command itself, users can use `%PROJECT_NAME%` to refer to the
project's name. This placeholder is replaced by the project name specified in the `gun init` command.
For example:
```console
$ gun init helloWorld py
```
In this case, the following config:
```toml
[init.py]
command = "echo %PROJECT_NAME% > README"
```
Would execute the command:
```console
echo helloWorld > README
```


## Project Language

Similarly, users can use `%PROJECT_LANGUAGE%` to refer to the project's language specified in the
gun init command. For instance, if a user initialises a project with the following command:

```console
$ gun init myProject java
```
The corresponding config:
```toml
[init.java]
command = "echo %PROJECT_LANGUAGE% > LANGUAGE.txt"
```
Would execute the command:
```console
echo java > LANGUAGE.txt
```

