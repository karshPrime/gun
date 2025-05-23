
// actions/trigger.go

package actions

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"karshPrime/gun/config"
)

//- Defines ----------------------------------------------------------------------------------------

const helpGlobal = "Run the command with globally defined config, override local config";

type triggerConfigs struct {
	global	bool;
	cdRoot	bool;
	command string;
};


//- Private Helpers --------------------------------------------------------------------------------

func projectLanguage() string {
	var lExtension string

	lCheckFiles := func ( aPattern string ) bool {
		lFiles, err := filepath.Glob( aPattern )
		if err != nil {
			fmt.Println( "Error while searching for files:", err )
			return false
		}

		for _, file := range lFiles {
			if lInfo, err := os.Stat(file); err == nil && !lInfo.IsDir() {
				lExtension = strings.ToLower( filepath.Ext(file) )
				return true
			}
		}
		return false
	}

	lFileFound := lCheckFiles( "main.*" ) || lCheckFiles( "app.*" )

	if lFileFound {
		return lExtension
	}

	fmt.Println( "Unable to find project language" );
	os.Exit( 1 );
	return ""
}

func ( configs *triggerConfigs ) parseInput() {
	lLastElement := len( os.Args ) -1;

	if lLastElement == 0 {
		return;
	}

	lIsGlobal := func ( aIndex int ) bool {
		return os.Args[aIndex] == "--global" || os.Args[aIndex] == "-g";
	}

	if lIsGlobal( 1 ) {
		configs.global = true;
		configs.command = strings.Join( os.Args[2:], " " )

	} else if lIsGlobal( lLastElement ) {
		configs.global = true;
		configs.command = strings.Join( os.Args[1:lLastElement], " " )

	} else {
		configs.command = strings.Join( os.Args[1:], " " )
	}
}

func ( configs *triggerConfigs ) parseConfigs( aCommand Triggers ) {
	if configs.global {
		lConfigFile := config.ConfigDir() + "config.toml";
		lProjectLanguage := projectLanguage();

		return;
	}
}


//- Public Calls -----------------------------------------------------------------------------------

func Trigger( aCommand Triggers ) {
	var lConfigs triggerConfigs;

	lConfigs.parseInput();
	lConfigs.parseConfigs( aCommand );

	if lConfigs.cdRoot {
		cdRoot();
	}

	// run command
	// lTriggerConfigs.command
}

