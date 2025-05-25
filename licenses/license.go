
// license/license.go

package licenses

import (
	"flag"
)

//- Defines ----------------------------------------------------------------------------------------

const helpHere = "Create project in current directory, instead of mkdir"
const helpLicense = "State what license to use"


//- Private Helpers --------------------------------------------------------------------------------

type licenseConfigs struct {}

func ( configs *licenseConfigs ) parseInput() {

	flag.Parse();
}

func ( configs *licenseConfigs ) parseConfigs() {
	//
}


//- Public Calls -----------------------------------------------------------------------------------

func License() {
	var lConfigs licenseConfigs;

	// arg 2 & 3 should be project name and language

	lConfigs.parseInput();
	lConfigs.parseConfigs();

	// run commands
}

