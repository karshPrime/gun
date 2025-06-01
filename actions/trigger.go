// actions/trigger.go

package actions

import ( 
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"karshPrime/gun/config"
	"karshPrime/gun/logs"
	"github.com/pelletier/go-toml"
 )

//- Defines ----------------------------------------------------------------------------------------

type triggerConfigs struct {
	global	bool;
	cdRoot	bool;
	command string;
};


//- Private Helpers --------------------------------------------------------------------------------

func projectLanguage() string {
	var lExtension string;
	var lFound bool = false;

	lCheckMatch := func( aName string ) bool {
		lBase := filepath.Base( aName );
		lMatchMain, _ := filepath.Match( "main.*", lBase );
		lMatchApp, _ := filepath.Match( "app.*", lBase );
		return lMatchMain || lMatchApp;
	}

	err := filepath.Walk( ".", func( aPath string, aInfo os.FileInfo, aErr error ) error {
		if aErr != nil {
			return aErr;
		}

		// skip if it's a directory
		if aInfo.IsDir() {
			lRelPath := strings.Count( filepath.Clean( aPath ), string( os.PathSeparator ))
			if lRelPath > 2 {
				return filepath.SkipDir;
			}
			return nil;
		}

		lDepth := strings.Count( filepath.Clean( aPath ), string( os.PathSeparator ))
		if lDepth > 2 {
			return nil;
		}

		if lCheckMatch( aPath ) {
			lExtension = strings.ToLower( filepath.Ext( aPath ))
			lFound = true;
			return filepath.SkipDir; // stop search early
		}
		return nil
	} )

	if err != nil {
		logs.ErrorPrint( "Unable to search for files in the directory" );
		fmt.Println( err );
		os.Exit( 1 );
	}

	if lFound {
		return lExtension;
	}

	logs.ErrorPrint( "Unable to find project language" );
	os.Exit( 1 );
	return "";
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
		configs.command = strings.Join( os.Args[2:], " " );

	} else if lIsGlobal( lLastElement ) {
		configs.global = true;
		configs.command = strings.Join( os.Args[1:lLastElement], " " );

	} else {
		configs.command = strings.Join( os.Args[1:], " " );
	}
}

func ( configs *triggerConfigs ) globalConfigParse( aTree *toml.Tree, aTrigger Triggers ) bool {
	lStatus := false;
	lTriggerKey := triggersKey( aTrigger );
	lProjectLanguage := projectLanguage();
	lLangKey := fmt.Sprintf( "dev.%s", strings.TrimPrefix( lProjectLanguage, "." ));

	section := aTree.Get( lLangKey );
	if section == nil {
		logs.ErrorPrint( lProjectLanguage + "config not found" );
		return false;
	}

	lSectionMap, ok := section.( *toml.Tree );
	if !ok {
		return false;
	}

	lCommand := lSectionMap.Get( lTriggerKey );
	if lCommandStr, ok := lCommand.( string ); ok {
		configs.command = lCommandStr + " " + configs.command;
		lStatus = true;
	}

	configs.cdRoot = false;
	if lcdRootValue := lSectionMap.Get( "cd_root" ); lcdRootValue != nil {
		configs.cdRoot, _ = lcdRootValue.( bool );
	}

	return lStatus;
}

func ( configs *triggerConfigs ) localConfigParse( aTrigger Triggers, aData string ) bool {
	lTriggerKey := triggersKey( aTrigger );

	lLines := strings.Split( string( aData ), "\n" );
	lCommandMap := make( map[string]string );

	for _, line := range lLines {
		line = strings.TrimSpace( line );
		if line == "" || strings.HasPrefix( line, "//" ) {
			continue;
		}

		if strings.Contains( line, "=" ) {
			lParts := strings.SplitN( line, "=", 2 );
			lKey := strings.TrimSpace( lParts[0] );
			lVal := strings.TrimSpace( lParts[1] );
			lVal = strings.Trim( lVal, `',` ); // remove quotes and commas

			if lKey == "cd_root" {
				configs.cdRoot = lVal == "true";
			} else {
				lCommandMap[lKey] = lVal;
			}
		}
	}

	lCommand, lStatus := lCommandMap[lTriggerKey];
	configs.command = lCommand + " " + configs.command;

	return lStatus;
}

func ( configs *triggerConfigs ) parseConfigs( aTrigger Triggers ) bool {
	lGlobalConfigFile := config.ConfigDir() + "config.toml";
	lGlobalConfigData, err := os.ReadFile( lGlobalConfigFile );
	if err != nil {
		return false;
	}

	lTree, err := toml.Load( string( lGlobalConfigData ));
	if err != nil {
		return false;
	}

	if configs.global {
		return configs.globalConfigParse( lTree, aTrigger );
	}

	lLocalConfigFile := "";
	if lLocalSection := lTree.Get( "local" ); lLocalSection != nil {
		if lLocalMap, ok := lLocalSection.( *toml.Tree ); ok {
			if lFileName := lLocalMap.Get( "config_title" ); lFileName != nil {
				lLocalConfigFile, _ = lFileName.( string );
			}
		}
	}

	if lLocalConfigFile == "" {
		lLocalConfigFile = "commands";
	}

	lData, err := os.ReadFile( lLocalConfigFile );
	if err != nil {
		// use global config if local commands file cant be found
		return configs.globalConfigParse( lTree, aTrigger );
	}

	return configs.localConfigParse( aTrigger, string( lData ));
}


//- Public Calls -----------------------------------------------------------------------------------

func Trigger( aCommand Triggers ) {
	var lConfigs triggerConfigs;

	lConfigs.parseInput();
	lStatus := lConfigs.parseConfigs( aCommand );

	if !lStatus { return; }

	if lConfigs.cdRoot {
		if !cdRoot() { return; };
	}

	// run command
	lResult, lError := sysRun( lConfigs.command );
	if lError {
		logs.ErrorPrint( lResult );
		return;
	}
	fmt.Print( lResult );
}

