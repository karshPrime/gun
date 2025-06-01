// go:build windows

// actions/lib_windows.go

package actions

func systemShell( aCommand string ) ( string, []string ) {
	lShell := "sh";
	lArgs  := []string{ "-c", aCommand };

	return lShell, lArgs;
}

