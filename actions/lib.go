
// actions/lib.go

package actions

//- Defines ----------------------------------------------------------------------------------------

type Triggers int;

const (
	Help Triggers = iota;
	Run; Build; Test; Debug; Clean;
)


//- Private Helpers --------------------------------------------------------------------------------

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

