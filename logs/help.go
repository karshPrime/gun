package logs

import (
	"os"
	"fmt"
	"github.com/fatih/color"
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

const aboutFlagLocal    = "Edit the local config file instead of the global one"
const aboutFlagGlobal   = "Run the command using the global config, overriding the local config";
const aboutFlagNoGit    = "Create the project but not git repository";
const aboutFlagHere     = "Create the project in the current dir, instead of making a new one";
const aboutFlagLicense  = "Use the specified license for the project instead of the config defined";
const aboutFlagList     = "List all saved ";  // + license / template
const aboutFlagAdd      = "Add a saved template to the current project";
const aboutFlagNew      = "Create a new "; // + license / template
const aboutFlagReplace  = "Replace the current license with a new one";
const aboutFlagManage   = "Manage all saved templates";
const aboutFlagPrintDir = "Directory where all files are saved";
const aboutFlagNoTemplates = "Do not copy any config defined templates to the project";

type printInfo struct {
	about	string;
	command	string;
};

//- Private Helpers --------------------------------------------------------------------------------

func helpAll() {
	lRed := color.New( color.FgHiRed ).SprintFunc();
	lBlue := color.New( color.FgHiBlue ).SprintFunc();
	lGray := color.New( color.FgHiBlack ).SprintFunc();
	lGreen := color.New( color.FgHiGreen ).SprintFunc();
	lYellow := color.New( color.FgHiYellow ).SprintFunc();
	lDullGreen := color.New( color.FgGreen ).SprintFunc();

	fmt.Println(
		"\nUSAGE:",
		"\n  ", lYellow("$"), lRed("gun"), lGreen( "[command]" ), lBlue( "[flags] [arguments]" ),
		"\n",
		"\nCOMMANDS:",
		lGreen( "\n  run\t\t" )   , lDullGreen( aboutCommandRun ),
		lGreen( "\n  build\t\t" ) , lDullGreen( aboutCommandBuild ),
		lGreen( "\n  init\t\t" )  , lDullGreen( aboutCommandInit ),
		lGreen( "\n  debug\t\t" ) , lDullGreen( aboutCommandDebug ),
		lGreen( "\n  test\t\t" )  , lDullGreen( aboutCommandTest ),
		lGreen( "\n  clean\t\t" ) , lDullGreen( aboutCommandClean ),
		lGreen( "\n  config\t" )  , lDullGreen( aboutCommandConfig ),
		lGreen( "\n  template\t" ), lDullGreen( aboutCommandTemplate ),
		lGreen( "\n  license\t" ) , lDullGreen( aboutCommandLicense ),
		"\n",
		"\nEXAMPLES:",
		lGray( "\n  # Compile and run the program using local config" ),
		lYellow( "\n  $" ), lRed( "gun" ),
		"\n",
		lGray( "\n  # Run the program with '--help' argument" ),
		lYellow( "\n  $" ), lRed( "gun" ), lGreen( "run" ), lBlue( "--help" ),
		"\n",
		lGray( "\n  # Compile with extra flags and run with arguments" ),
		lYellow( "\n  $" ), lRed( "gun" ), lBlue( "\"hello world\" --flags -Wall" ),
		"\n",
		lGray( "\n  # Initialise a new C project using GPLv2 license" ),
		lYellow( "\n  $" ), lRed( "gun" ), lGreen( "init" ), lBlue( "myProject c --license GPLv2" ),
		"\n",
		lGray( "\n  # List all available templates" ),
		lYellow( "\n  $" ), lRed( "gun" ), lGreen( "template" ), lBlue( "--list" ),
		"\n",
		"\nNOTE:",
		lGray( "\n  Run `$ gun help [command]` to see details, flags, and usage for a specific" ),
		lGray( "command.\n  For example: `$ gun help build`" ),
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
	lRed := color.New( color.FgHiRed ).SprintFunc();
	lBlue := color.New( color.FgHiBlue ).SprintFunc();
	lGray := color.New( color.FgHiBlack ).SprintFunc();
	lGreen := color.New( color.FgHiGreen ).SprintFunc();
	lYellow := color.New( color.FgHiYellow ).SprintFunc();
	lDullGreen := color.New( color.FgGreen ).SprintFunc();

	var lFlags []printInfo;
	var lExamples []printInfo;
	var lParameters string = "[arguments] [flags]";
	var lNote string;

	switch ( aCommand ) {
		case "run":
			fmt.Println( aboutCommandRun );
			lParameters = "[arguments]";
			lExamples = []printInfo{{
				about: "Run the last compiled Go program",
				command: "",
			},{
				about: "Run the last compiled C program",
				command: "",
			},{
				about: "Run the last compiled program with `\"hello world\" 1234` passed as arguments",
				command: "\"hello world\" 1234",
			}};

		case "build":
			fmt.Println( aboutCommandBuild );
			lFlags = []printInfo{
				{ about: aboutFlagGlobal, command: "global" },
			};
			lExamples = []printInfo{{
				about: "Build the Go project",
				command: "",
			},{
				about: "Build the Java project with `-g:none` parameters",
				command: "-g:none",
			},{
				about: "Build the C flag with default configs and with -Wextra -std=c11 parameters",
				command: "-Wextra -std=c11 --global",
			}};

		case "debug":
			fmt.Println( aboutCommandDebug );
			lFlags = []printInfo{
				{ about: aboutFlagGlobal, command: "global" },
			};
			lExamples = []printInfo{{
				about: "Run the set debugger for the project",
				command: "",
			},{
				about: "Run the default debugger with `-batch -ex \"run\" -ex \"bt\"` arguments",
				command: "-batch -ex \"run\" -ex \"bt\" --global",
			}};

		case "test":
			fmt.Println( aboutCommandTest );
			lFlags = []printInfo{
				{ about: aboutFlagGlobal, command: "global" },
			};
			lExamples = []printInfo{{
				about: "Run the defined test suite",
				command: "",
			},{
				about: "Run the global defined test suit with `\"foo\" \"bar\"` arguments",
				command: "foo bar --global",
			}};

		case "clean":
			fmt.Println( aboutCommandClean );
			lFlags = []printInfo{
				{ about: aboutFlagGlobal, command: "global" },
			};
			lExamples = []printInfo{{
				about: "Clean the project",
				command: "",
			},{
				about: "Clean the project as defined in global config",
				command: "--global",
			}};
			lNote = "You can create shell scripts for each project to define additional"+
			" subcommands, such as\n  `fullclean` and buildclean by including `sh" +
			" cleanscript.sh` in your config.\n  Alternatively, you can simply add a command like" +
			" \"rm -rf ./build/*\"";


		case "init":
			fmt.Println( aboutCommandInit );
			lFlags = []printInfo{
				{ about: aboutFlagHere,        command: "here" },
				{ about: aboutFlagNoGit,       command: "no-git" },
				{ about: aboutFlagNoTemplates, command: "no-templates" },
				{ about: aboutFlagLicense,     command: "license [title]", },
			};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};
			lParameters = "<project name> <project language> [flags]";

		case "template":
			fmt.Println( aboutCommandTemplate );
			lFlags = []printInfo{
				{ about: aboutFlagList+"template", command: "list" },
				{ about: aboutFlagPrintDir, command:"print-dir" },
				{ about: aboutFlagNew+"template", command: "new [title]" },
				{ about: aboutFlagAdd,      command: "add [title]" },
				{ about: aboutFlagManage,   command: "manage"},
			};
			lExamples = []printInfo{{
				about: "",
				command: "",
			}};

		case "config":
			fmt.Println( aboutCommandConfig );
			lFlags = []printInfo{
				{ about: aboutFlagLocal, command: "local" },
			};
			lExamples = []printInfo{{
				about: "Update global config",
				command: "",
			},{
				about: "Update local config",
				command: "--local",
			}};

		case "license":
			fmt.Println( aboutCommandLicense );
			lFlags = []printInfo{
				{ about: aboutFlagList+"license",     command: "list" },
				{ about: aboutFlagPrintDir, command: "print-dir" },
				{ about: aboutFlagNew+"license",      command: "new [title]" },
				{ about: aboutFlagReplace,  command: "replace [title]"},
			};
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
		lYellow( "\n  $" ), lRed( "gun" ), lGreen( aCommand ), lBlue( lParameters ),
	);

	if len( lFlags ) > 0 {
		fmt.Println( "\nFLAGS:" );

		for i := range len( lFlags ) {
			fmt.Printf( 
				"  %-27v %v\n",
				lGreen( "--", lFlags[i].command ), lDullGreen( lFlags[i].about ),
			);
		}
	}

	fmt.Print( "\nEXAMPLES:" );

	for i := range len( lExamples ) {
		fmt.Println(
			lGray( "\n  # ", lExamples[i].about, "\n" ),
			lYellow( " $" ), lRed( "gun" ), lGreen( aCommand ), lBlue( lExamples[i].command ),
		);
	}

	if len(lNote) > 0 {
		fmt.Println( "\nNote:\n ",  lGray( lNote ) );
	}
}

