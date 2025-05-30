
package main

import (
	"os"
	"karshPrime/gun/actions"
	"karshPrime/gun/config"
	"karshPrime/gun/licenses"
	"karshPrime/gun/logs"
	"karshPrime/gun/templates"
)

func main() {
	config.ValidateFilesystem();

	if len( os.Args ) < 2 {
		actions.BuildRun();
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
		case "h", "help", "--help", "-h" : logs.Help();

		case "T", "template" : templates.Template();
		case "l", "license"  : licenses.License();

		default:
			actions.BuildRun();
	}
}

