
// config/lib.go

package config

import (
	"os"
	"fmt"
)


//- Defines ----------------------------------------------------------------------------------------

//- Private Helpers --------------------------------------------------------------------------------


//- Public Calls -----------------------------------------------------------------------------------

func ConfigDir() string  {
	lHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err );
        os.Exit( 1 );
    }

	return lHomeDir + "/.config/gun/";
}

