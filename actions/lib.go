
// actions/lib.go

package actions

import (
	"io"
	"os"
	"bytes"
	"strings"
	"os/exec"
	"path/filepath"
	"karshPrime/gun/logs"
)

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

func cdRoot() bool {
	lResult, lError := SysRun( "git rev-parse --show-toplevel" );
	if lError {
		logs.ErrorPrint( "Project is not a git repo. Cannot cd to project root.\n", lResult );
		return false;
	}

	lErrorCD := os.Chdir( lResult );
    if lErrorCD != nil {
        logs.ErrorPrint( "Unable to cd to project root.\n", lErrorCD )
		return false;
    }

	return true;
}

func copyDir( aSource string, aDestination string ) ( fail bool ) {
	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll( aDestination, os.ModePerm ); err != nil {
		logs.ErrorPrint( "Failed to create directory %s: %w", aDestination, err );
		return true;
	}

	// Read the contents of the source directory
	entries, err := os.ReadDir( aSource );
	if err != nil {
		logs.ErrorPrint( "Failed to read directory %s: %w", aSource, err );
		return true;
	}

	// Copy each entry in the directory
	for _, entry := range entries {
		srcPath := filepath.Join( aSource, entry.Name() );
		dstPath := filepath.Join( aDestination, entry.Name() );

		if entry.IsDir() {
			// Recursively copy subdirectories
			if copyDir( srcPath, dstPath ) {
				return true;
			}
		} else {
			// Copy files
			if copyFile( srcPath, dstPath ) {
				return true;
			}
		}
	}

	return false;
}

func copyFile( aSource string, aDestination string ) ( fail bool ) {
	// Create the destination directory if it doesn't exist
	dstDir := filepath.Dir( aDestination );
	if err := os.MkdirAll( dstDir, os.ModePerm ); err != nil {
		logs.ErrorPrint( "Failed to create directory %s: %w", dstDir, err );
		return false;
	}

	// Open the source file
	sourceFile, err := os.Open( aSource );
	if err != nil {
		logs.ErrorPrint( "failed to open source file %s: %w", aSource, err );
		return true;
	}
	defer sourceFile.Close();

	// Create the destination file
	destinationFile, err := os.Create( aDestination );
	if err != nil {
		logs.ErrorPrint( "failed to create destination file %s: %w", aDestination, err );
		return true;
	}
	defer destinationFile.Close()

	// Copy the contents from source to destination
	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		logs.ErrorPrint(
			"failed to copy file contents from %s to %s: %w", aSource, aDestination, err,
		);
		return true;
	}

	return false;
}


//- Public Calls -----------------------------------------------------------------------------------

func SysRun( aCommand string ) ( Result string, Error bool ) {
	logs.DebugPrint( aCommand );
	lShell, lArgs := systemShell( aCommand );

	if lShell == "" {
		return "invalid environment", true;
	}

	var lStdOut bytes.Buffer;
	var lStdErr bytes.Buffer;

	lCommand := exec.Command( lShell, lArgs... );
	lCommand.Stdout = &lStdOut;
	lCommand.Stderr = &lStdErr;

	if lCommand.Run() != nil {
		return strings.TrimSuffix( lStdErr.String(), "\n" ) +
			strings.TrimSuffix( lStdOut.String(), "\n" ), true;
	}

	return strings.TrimSuffix( lStdErr.String(), "\n" ) +
		strings.TrimSuffix( lStdOut.String(), "\n" ), false;
}

func Copy( aSource string, aDestination string) ( fail bool ) {
	// Check if the source is a directory
	lSrcInfo, err := os.Stat( aSource )
	if err != nil {
		logs.ErrorPrint( "Failed to stat source %s: %w", aSource, err );
		return true;
	}

	if lSrcInfo.IsDir() {
		return copyDir( aSource, aDestination );
	} else {
		return copyFile( aSource, aDestination );
	}
}

