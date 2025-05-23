
// actions/trigger.go

package actions

import (
	"fmt"
	"os"
	"strings"
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
	lLastElement := len( os.Args ) -1;

	if lLastElement == 0 {
		return;
	}

	lIsGlobal := func ( aIndex int ) bool {
		return os.Args[aIndex] == "--global" || os.Args[aIndex] == "-g";
	}

	if lIsGlobal( 1 ) {
		configs.global = true;
		configs.command = strings.Join( os.Args[2:], " " )

	} else if lIsGlobal( lLastElement ) {
		configs.global = true;
		configs.command = strings.Join( os.Args[1:lLastElement], " " )

	} else {
		configs.command = strings.Join( os.Args[1:], " " )
	}
}

func ( configs *triggerConfigs ) parseConfigs( aCommand Triggers ) {
	if configs.global {
		//

		return;
	}
}


//- Public Calls -----------------------------------------------------------------------------------

func Trigger( aCommand Triggers ) {
	var lConfigs triggerConfigs;

	lConfigs.parseInput();
	lConfigs.parseConfigs( aCommand );

	if lConfigs.cdRoot {
		cdRoot();
	}

	// run command
	// lTriggerConfigs.command
}

