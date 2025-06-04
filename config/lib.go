
// config/lib.go

package config

import (
	"os"
	"fmt"
)


//- Defines ----------------------------------------------------------------------------------------

const GunURL = "https://github.com/karshPrime/gun";

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

