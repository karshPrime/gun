
// actions/lib_unix.go
// go:build linux,darwin

package actions

func systemShell( aCommand string ) ( string, []string ) {
	lShell := "sh";
	lArgs  := []string{ "-c", aCommand };

	return lShell, lArgs;
}

