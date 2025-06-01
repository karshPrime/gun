// +build NDEBUG

// logs/print.go

package logs

import (
	"os"
	"fmt"
	"github.com/fatih/color"
)

//- Public Calls -----------------------------------------------------------------------------------

func ErrorPrint( aPrompt string ) {
	lRed := color.New( color.FgHiRed ).SprintFunc();

	fmt.Fprintln( os.Stderr, lRed("[Error]"), aPrompt );
}

func WarningPrint( aPrompt string ) {
	lYellow := color.New( color.FgHiYellow ).SprintFunc();

	fmt.Fprintln( os.Stderr, lYellow("[Warning]"), aPrompt );
}

func DebugPrint( _ string ) {
	return;
}

