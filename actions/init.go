
// actions/init.go

package actions

import (
	"flag"
)

//- Defines ----------------------------------------------------------------------------------------

const helpHere = "Create project in current directory, instead of mkdir"
const helpLicense = "State what license to use"

type initConfigs struct {
	here		bool;
	license		string;
	noGit		bool;
	gitIgnore		[]string;
	gitOnlyIgnore	[]string;
	template		[]string;
	ignoreTemplates bool;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *initConfigs ) parseInput() {
	flag.BoolVar( &configs.here, "here", false, helpHere );
	flag.BoolVar( &configs.noGit, "no-git", false, helpHere );
	flag.BoolVar( &configs.ignoreTemplates, "ignore-template", false, helpHere );
	flag.StringVar( &configs.license, "license", "", helpLicense );

	// add for string arrays

	flag.Parse();
}

func ( configs *initConfigs ) parseConfigs() {
	//
}

//- Public Calls -----------------------------------------------------------------------------------

func Init() {
	var lInitConfigs initConfigs;

	// arg 2 & 3 should be project name and language

	lInitConfigs.parseInput();
	lInitConfigs.parseConfigs();

	// run commands
}

