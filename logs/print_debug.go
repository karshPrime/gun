// +build !NDEBUG

// logs/print.go

package logs

import (
	"os"
	"fmt"
	"runtime"
	"github.com/fatih/color"
)

//- Private Helpers --------------------------------------------------------------------------------

func lineInfo() string  {
	_, lFile, lLine, ok := runtime.Caller( 2 )
	if ok {
		return fmt.Sprintf( "(%s:%d)", lFile, lLine )
	} else {
		return "(Could not retrieve caller information)"
	}
}

//- Public Calls -----------------------------------------------------------------------------------

func ErrorPrint( aPrompt string ) {
	lRed := color.New( color.FgRed ).SprintFunc();
	lLiRed := color.New( color.FgHiRed ).SprintFunc();

	fmt.Fprintln( os.Stderr, lRed("[ERROR]"), lLiRed(lineInfo()), aPrompt );
}

func WarningPrint( aPrompt string ) {
	lYellow := color.New( color.FgYellow ).SprintFunc();
	lLiYellow := color.New( color.FgHiYellow ).SprintFunc();

	fmt.Fprintln( os.Stderr, lYellow("[WARNING]"), lLiYellow(lineInfo()), aPrompt );
}

func DebugPrint( aPrompt string ) {
	lBlue := color.New( color.FgBlue ).SprintFunc();
	lLiBlue := color.New( color.FgHiBlue ).SprintFunc();

	fmt.Fprintln( os.Stderr, lBlue("[DEBUG]"), lLiBlue(lineInfo()), aPrompt );
}

