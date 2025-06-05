// main.go

package main

//- Imports ----------------------------------------------------------------------------------------

import (
	"os"
	"fmt"
	"karshPrime/gun/actions"
	"karshPrime/gun/config"
	"karshPrime/gun/licenses"
	"karshPrime/gun/logs"
	"karshPrime/gun/templates"
)

const VERSION = "0.1.0";

//- Main -------------------------------------------------------------------------------------------

func main() {
	config.ValidateFilesystem();

	if len( os.Args ) < 2 {
		actions.BuildRun( os.Args );
		return;
	}

	lOriginalArgs := os.Args;
	os.Args = append( []string{ lOriginalArgs[0] }, lOriginalArgs[2:]... );

	switch ( lOriginalArgs[1] ) {
		case "r", "run"   : actions.Trigger( actions.RUN   );
		case "b", "build" : actions.Trigger( actions.BUILD );
		case "c", "clean" : actions.Trigger( actions.CLEAN );
		case "d", "debug" : actions.Trigger( actions.DEBUG );
		case "t", "test"  : actions.Trigger( actions.TEST  );

		case "i", "init" : actions.Init();

		case "T", "template" : templates.Template();
		case "l", "license"  : licenses.License();

		case "h", "help", "--help", "-h" 	   : logs.Help();
		case "v", "version", "--version", "-v" : fmt.Println( "Version:", VERSION );

		default:
			actions.BuildRun( lOriginalArgs );
	}
}

