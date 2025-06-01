
// actions/lib_windows.go
// go:build windows

package actions

func systemShell( aCommand string ) ( string, []string ) {
	lShell := "sh";
	lArgs  := []string{ "-c", aCommand };

	return lShell, lArgs;
}

