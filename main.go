
package main

import (
	"os"
	"karshPrime/gun/actions"
	"karshPrime/gun/licenses"
	"karshPrime/gun/logs"
	"karshPrime/gun/templates"
)

func main() {
    if len( os.Args ) < 2 {
		actions.BuildRun();
		return;
    }

	lOriginalArgs := os.Args;
	os.Args = append( []string{ lOriginalArgs[0] }, lOriginalArgs[2:]... );

	switch ( lOriginalArgs[1] ) {
		case "r", "run"   : actions.Trigger( actions.Run   );
		case "b", "build" : actions.Trigger( actions.Build );
		case "c", "clean" : actions.Trigger( actions.Clean );
		case "d", "debug" : actions.Trigger( actions.Debug );
		case "t", "test"  : actions.Trigger( actions.Test  );

		case "i", "init" : actions.Init();
		case "h", "help" : logs.Help();

		case "T", "template" : templates.Template();
		case "l", "license"  : licenses.License();

		default:
			actions.BuildRun();
	}
}

