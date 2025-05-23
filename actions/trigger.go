// actions/trigger.go

package actions

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"karshPrime/gun/config"
	"github.com/pelletier/go-toml"
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

		lConfigData, err := os.ReadFile( lConfigFile )
		if err != nil {
			return;
		}

		lTree, err := toml.Load( string(lConfigData) )
		if err != nil {
			return;
		}

		lLangKey := fmt.Sprintf( "dev.%s", strings.TrimPrefix( lProjectLanguage, "." ))

		section := lTree.Get( lLangKey )
		if section == nil {
			return;
		}

		lSectionMap, ok := section.( *toml.Tree )
		if !ok {
			return;
		}

		lCommand := lSectionMap.Get( triggersKey( aCommand ))

		if lCommandStr, ok := lCommand.( string ); ok {
			configs.command = lCommandStr + " " + configs.command
		}

		configs.cdRoot = false
		if lcdRootValue := lSectionMap.Get("cd_root"); lcdRootValue != nil {
			configs.cdRoot, _ = lcdRootValue.( bool )
		}
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

