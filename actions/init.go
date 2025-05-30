// actions/init.go

package actions

import (
	"flag"
	"karshPrime/gun/logs"
	"os"
)

//- Defines ----------------------------------------------------------------------------------------

const helpHere = "Create project in current directory, instead of mkdir"
const helpLicense = "State what license to use"

type initConfigs struct {
	here		bool;
	noGit		bool;
	noTemplates bool;
	license		string;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *initConfigs ) parseInput() {
	flag.BoolVar( &configs.here, "here", false, helpHere );
	flag.BoolVar( &configs.noGit, "no-git", false, helpHere );
	flag.BoolVar( &configs.noTemplates, "ignore-templates", false, helpHere );
	flag.StringVar( &configs.license, "license", "", helpLicense );

	flag.Parse();
}

func ( configs *initConfigs ) parseConfigs( aProjectLanguage string ) {
	//
}

//- Public Calls -----------------------------------------------------------------------------------

func Init() {
	var lConfigs initConfigs;

	if len( os.Args ) < 3 {
		// errorMissingArgs();
		logs.HelpCommand("init");
		return;
    }

	lOriginalArgs := os.Args;
	os.Args = append( []string{ lOriginalArgs[0] }, lOriginalArgs[3:]... );

	// lOriginalArgs[1] = project name
	// lOriginalArgs[2] = project language

	lConfigs.parseInput();
	lConfigs.parseConfigs( lOriginalArgs[2] );

	// run commands
	if !lConfigs.here {
		// mkdir lOriginalArgs[1];
		// cd lOriginalArgs[1]
	}

	if !lConfigs.noTemplates {
		// for template in lConfigs.templates cp
	}

	if !lConfigs.noGit {
		// git init
		// cp CONFIG_DIR/ignores/language ./.gitginore
		// git commit init: project
	}
}

