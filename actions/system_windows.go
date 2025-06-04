// go:build windows

// actions/lib_windows.go

package actions

func systemShell( aCommand string ) ( string, []string ) {
	lShell := "cmd.exe";
	lArgs  := []string{ "/C", aCommand };
	return lShell, lArgs;
}

