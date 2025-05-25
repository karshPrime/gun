
// template/template.go

package templates

import (
	"flag"
)

//- Defines ----------------------------------------------------------------------------------------


//- Private Helpers --------------------------------------------------------------------------------

type templateConfigs struct {}

func ( configs *templateConfigs ) parseInput() {

	flag.Parse();
}

func ( configs *templateConfigs ) parseConfigs() {
	//
}


//- Public Calls -----------------------------------------------------------------------------------

func Template() {
	var lConfigs templateConfigs;

	lConfigs.parseInput();
	lConfigs.parseConfigs();

	// run commands
}

