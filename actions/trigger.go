
// actions/trigger.go

package actions

import (
	"flag"
)

//- Defines ----------------------------------------------------------------------------------------

const helpGlobal = "the program with globally defined config, override local config";

type triggerConfigs struct {
	global	bool;
	cdRoot	bool;
	command string;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *triggerConfigs ) parseInput( aCommand string ) {
	flag.BoolVar( &configs.global, "global", false, aCommand + helpGlobal );

	flag.Parse();
}

func ( configs *triggerConfigs ) parseConfigs() {
	if configs.global {
		//

		return;
	}
}


//- Public Calls -----------------------------------------------------------------------------------

func Trigger( aCommand string ) {
	var lTriggerConfigs triggerConfigs;

	lTriggerConfigs.parseInput( aCommand );
	lTriggerConfigs.parseConfigs();

	if lTriggerConfigs.cdRoot {
		cdRoot();
	}

	// run command
	// lTriggerConfigs.command
}

