
$ gun
    --flags [Flags]                     Specify optional build flags
    --args  [Args]                      Specify optional run arguments
    --global                            Build and Run the program with globally defined config, override local config

$ gun init [ProjectName] [Language]
    --here                              Create project in current directory, instead of mkdir
    --license [Title]                   State what license to use
    --no-git                            Not initialise as git repository
    --no-templates                      Dont add any or all template to the project

$ gun template
    --list                              List all template files
    --print-dir                         Print the location of where the templates are stored
    --new [Title]                       Create new template file with the given title
    --add [Templates]                   Add templates code to current project
    --manage                            Manage the record of all saved templates

$ gun license
    --list                              List all saved licenses
    --print-dir                         Print the location of where the licenses are stored
    --new [Title]                       Create new license file with the given title
    --replace [Title]                   Replace the current license with new one

$ gun run {Args}                        Run the program with passed optional arguments
    --global                            Run the program with globally defined config, override local config

$ gun build {Flags}                   Compile the program with passed optional flags
    --global                            Compile the program with globally defined config, override local config

$ gun clean                             Clean the build files
    --global                            Clean the program with globally defined config, override local config

$ gun debug                             Run project specific debugger
    --global                            Run the debugger defined in the with global config, override local config

$ gun test                              Run the test files
    --global                            Run the test defined in the with global config, override local config

$ gun help {command}                    Print this help menu

$ gun version                           Print utility version

$ gun config
    --local                             Edit (if existing or create) local config file

