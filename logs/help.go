package logs

import (
	"fmt"
	"os"
)

//- Defines ----------------------------------------------------------------------------------------

const aboutCommandRun   = "Run the program with optional arguments";
const aboutCommandBuild = "Compile the program with optional flags";
const aboutCommandDebug = "Run the configured debugger for the project";
const aboutCommandTest  = "Run the test suite";
const aboutCommandClean = "Remove build artifacts";
const aboutCommandInit  = "Create a new project with templates, license, and structure";
const aboutCommandConfig   = "Edit or create a local project configuration file";
const aboutCommandTemplate = "Manage project templates (add, create, list)";
const aboutCommandLicense  = "Manage license files (create, replace, list)";

const aboutFlagGlobal = "";

type printInfo struct {
	about	string;
	command	string;
};

//- Private Helpers --------------------------------------------------------------------------------

func helpAll() {
	fmt.Println(
		"\n",
		"\nUSAGE:",
		"\n  $ gun [command] [flags]",
		"\n",
		"\nCOMMANDS:",
		"\n  run\t\t"   , aboutCommandRun,
		"\n  build\t\t" , aboutCommandBuild,
		"\n  init\t\t"  , aboutCommandInit,
		"\n  debug\t\t" , aboutCommandDebug,
		"\n  test\t\t"  , aboutCommandTest,
		"\n  clean\t\t" , aboutCommandClean,
		"\n  config\t"  , aboutCommandConfig,
		"\n  template\t", aboutCommandTemplate,
		"\n  license\t" , aboutCommandLicense,
		"\n",
		"\nEXAMPLES:",
		"\n  # Compile and run the program using local config",
		"\n  $ gun",
		"\n",
		"\n  # Run the program with '--help' argument",
		"\n  $ gun run --help",
		"\n",
		"\n  # Compile with extra flags and run with arguments",
		"\n  $ gun --flags -Wall --args \"hello world\"",
		"\n",
		"\n  # Initialise a new C project using GPLv2 license",
		"\n  $ gun init my_project c --license GPLv2",
		"\n",
		"\n  # List all available templates",
		"\n  $ gun template --list",
		"\n",
		"\nNOTE:",
		"\n  Run `$ gun help [command]` to see details, flags, and usage for a specific command.",
		"\n  For example: `$ gun help build`",
	);
}

//- Public Calls -----------------------------------------------------------------------------------

func Help() {
	if len( os.Args ) < 2 {
		fmt.Println( "Unify build/run/debug workflows across languages and project types." );
		helpAll();

		return;
	}

	HelpCommand( os.Args[1] );
}

func HelpCommand( aCommand string ) {
	var lFlags []printInfo;
	var lExamples []printInfo;
	var lParameters string = "";

	switch ( aCommand ) {
		case "run":
			fmt.Println( aboutCommandRun );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "global",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "build":
			fmt.Println( aboutCommandBuild );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "global",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "debug":
			fmt.Println( aboutCommandDebug );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "global",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "test":
			fmt.Println( aboutCommandTest );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "global",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "clean":
			fmt.Println( aboutCommandClean );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "global",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "init":
			fmt.Println( aboutCommandInit );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "here",
			},{
				about: aboutFlagGlobal,
				command: "license",
			},{
				about: aboutFlagGlobal,
				command: "no-git",
			},{
				about: aboutFlagGlobal,
				command: "no-templates",
			},{
				about: aboutFlagGlobal,
				command: "license",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "template":
			fmt.Println( aboutCommandTemplate );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "list",
			},{
				about: aboutFlagGlobal,
				command: "print-dir",
			},{
				about: aboutFlagGlobal,
				command: "new [title]",
			},{
				about: aboutFlagGlobal,
				command: "add [title]",
			},{
				about: aboutFlagGlobal,
				command: "manage",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "config":
			fmt.Println( aboutCommandConfig );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "local",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "license":
			fmt.Println( aboutCommandLicense );
			lFlags = []printInfo{{
				about: aboutFlagGlobal,
				command: "list",
			},{
				about: aboutFlagGlobal,
				command: "print-dir",
			},{
				about: aboutFlagGlobal,
				command: "new [title]",
			},{
				about: aboutFlagGlobal,
				command: "replace [title]",
			}};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		default:
			ErrorPrint( "invalid command " + aCommand );
			helpAll();
			return;
	}

	fmt.Println(
		"\nUSAGE:",
		"\n  $ gun", aCommand, lParameters,
		"\n",
		"\nFLAGS:",
	);

	for i := range len( lFlags ) {
		fmt.Printf( "  --%v\t\t%v\n", lFlags[i].command, lFlags[i].about );
	}

	fmt.Println( "\nEXAMPLES:" );

	for i := range len( lExamples ) {
		fmt.Printf(
			"  # %v\n  $ gun %v %v\n",
			lExamples[i].about , aCommand, lExamples[i].command,
		);
	}
}

