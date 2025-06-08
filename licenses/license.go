// license/license.go

package licenses

//- Imports ----------------------------------------------------------------------------------------

import (
	"os"
	"fmt"
	"flag"
	"karshPrime/gun/logs"
	"karshPrime/gun/config"
	"karshPrime/gun/actions"
)

//- Defines ----------------------------------------------------------------------------------------

type licenseConfigs struct {
	list		bool;
	printDir	bool;
	newLicense  string;
	replace		string;
}

//- Private Helpers --------------------------------------------------------------------------------

func ( configs *licenseConfigs ) parseInput() {
	flag.StringVar( &configs.newLicense, "new", "", "" );
	flag.StringVar( &configs.replace, "replace", "", "" );
	flag.BoolVar( &configs.printDir, "print-dir", false, "" )
	flag.BoolVar( &configs.list, "list", false, "" )

	flag.Usage = func() {
		logs.HelpCommand( "license", false );
	};

	flag.Parse();
}

func printLicenses( aDirectory string ) {
	// Read the directory
    files, err := os.ReadDir( aDirectory );
    if err != nil {
        logs.ErrorPrint( "Unable to read directory " + aDirectory, "\n", err );
    }

    // Loop through the files and print their names
    for _, file := range files {
        if !file.IsDir() { // check if it's not a directory
            fmt.Println( file.Name() )
        }
    }
}


//- Public Calls -----------------------------------------------------------------------------------

func License() {
	logs.DebugPrint( "called" );

	var lConfigs licenseConfigs;
	lConfigs.parseInput();

	lDirectory := config.ConfigDir() + "licenses";

	switch( true ) {
		case lConfigs.list:
			printLicenses( lDirectory );

		case lConfigs.printDir:
			fmt.Println( lDirectory );

		case lConfigs.newLicense != "":
			var lLicenseName string;

			if len(os.Args) < 4 {
				lLicenseName = os.Args[2];
			} else {
				lLicenseName = os.Args[3];
			}

			actions.Copy( lConfigs.newLicense, lDirectory + "/" + lLicenseName );

		case lConfigs.replace != "":
			actions.Copy( lDirectory + "/" + lConfigs.replace, "./LICENSE" );

		default:
			logs.ErrorPrint( "Invalid Usage" )
			logs.HelpCommand( "license", false );
	}
}

