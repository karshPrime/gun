
// actions/lib.go

package actions

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

func isGitRepo() bool {
	lIsRepo := false;

	// update lIsRepo

	return lIsRepo;
}

func cdRoot() {
	if isGitRepo() {
		// cd to git root
	}

	// print error
}

