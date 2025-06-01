
// actions/lib.go

package actions

import (
	"os"
	"bytes"
	"strings"
	"os/exec"
	"karshPrime/gun/logs"
)

//- Defines ----------------------------------------------------------------------------------------

type Triggers int;

const (
	HELP Triggers = iota;
	RUN; BUILD; TEST; DEBUG; CLEAN;
)


//- Private Helpers --------------------------------------------------------------------------------

func triggersKey( aTrigger Triggers ) string {
	switch aTrigger {
		case RUN   : return "run";
		case BUILD : return "build";
		case TEST  : return "test";
		case DEBUG : return "debug";
		case CLEAN : return "clean";
		default    : return "";
	}
}

func cdRoot() bool {
	lResult, lError := SysRun( "git rev-parse --show-toplevel" );
	if lError {
		logs.ErrorPrint( "Project is not a git repo. Cannot cd to project root.\n", lResult );
		return false;
	}

	lErrorCD := os.Chdir( lResult );
    if lErrorCD != nil {
        logs.ErrorPrint( "Unable to cd to project root.\n", lErrorCD )
		return false;
    }

	return true;
}

func SysRun( aCommand string ) ( Result string, Error bool ) {
	logs.DebugPrint( aCommand );
	lShell, lArgs := systemShell( aCommand );

	if lShell == "" {
		return "invalid environment", true;
	}

	var lStdOut bytes.Buffer;
	var lStdErr bytes.Buffer;

	lCommand := exec.Command( lShell, lArgs... );
	lCommand.Stdout = &lStdOut;
	lCommand.Stderr = &lStdErr;

	if lCommand.Run() != nil {
		return strings.TrimSpace( lStdErr.String() ), true;
	}

	return strings.TrimSpace( lStdOut.String() ), false;
}

