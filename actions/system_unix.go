// go:build linux,darwin

// actions/lib_unix.go

package actions

func systemShell( aCommand string ) ( string, []string ) {
	lShell := "sh";
	lArgs  := []string{ "-c", aCommand };

	return lShell, lArgs;
}

