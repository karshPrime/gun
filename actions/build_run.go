
// actions/build_run.go

package actions

import "os"

//- Public Calls -----------------------------------------------------------------------------------

func BuildRun( aOriginalArgs []string ) {
	lFoundFlags := false
	var lRunArgs, lBuildArgs []string = []string{""}, []string{""}

	for _, arg := range aOriginalArgs[1:] {
		if arg == "--flags" {
			lFoundFlags = true
		} else if lFoundFlags {
			lBuildArgs = append( lBuildArgs, arg )
		} else {
			lRunArgs = append( lRunArgs, arg )
		}
	}

	// set build args as os.Args
	os.Args = lBuildArgs;
	Trigger( BUILD );

	// set run flags as os.Args
	os.Args = lRunArgs;
	Trigger( RUN );
}

