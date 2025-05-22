
// actions/trigger.go

package actions

import (
	"fmt"
	"flag"
)

//- Defines ----------------------------------------------------------------------------------------

const helpGlobal = "Run the command with globally defined config, override local config";

type triggerConfigs struct {
	global	bool;
	cdRoot	bool;
	command string;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *triggerConfigs ) parseInput() {
	flag.BoolVar( &configs.global, "global", false, helpGlobal );

	flag.Parse();
}

func ( configs *triggerConfigs ) parseConfigs( aCommand Triggers ) {
	if configs.global {
		//

		return;
	}
}


//- Public Calls -----------------------------------------------------------------------------------

func Trigger( aCommand Triggers ) {
	fmt.Println( "trigger called with", aCommand )
	var lTriggerConfigs triggerConfigs;

	lTriggerConfigs.parseInput();
	lTriggerConfigs.parseConfigs( aCommand );

	if lTriggerConfigs.cdRoot {
		cdRoot();
	}

	// run command
	// lTriggerConfigs.command
}

