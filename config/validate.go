
package config

import (
	"os"
)


//- Defines ----------------------------------------------------------------------------------------


//- Private Helpers --------------------------------------------------------------------------------


//- Public Calls -----------------------------------------------------------------------------------

func ValidateFilesystem() {
	lHomeDir := ConfigDir();

	lCheckCreateDir := func( aDirPath string ) {
		if _, err := os.Stat( aDirPath ); os.IsNotExist( err ) {
			os.Mkdir( aDirPath, 0755 )
		}
	}

	lCheckCreateDir( lHomeDir );
	lCheckCreateDir( lHomeDir + "templates" )
	lCheckCreateDir( lHomeDir + "licenses" )
	lCheckCreateDir( lHomeDir + "scripts" )
}

