// actions/init.go

package actions

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"karshPrime/gun/logs"
	"karshPrime/gun/config"
	"github.com/pelletier/go-toml"
)

//- Defines ----------------------------------------------------------------------------------------

const helpHere = "Create project in current directory, instead of mkdir"
const helpLicense = "State what license to use"

type templates struct {
	title		string;
	destination	string;
}

type initConfigs struct {
	here		bool;
	command		string;
	license		string;
	noGit		bool;
	gitIgnores	[]string;
	noTemplates bool;
	templates	[]templates;
	directories	[]string;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *initConfigs ) parseInput() {
	flag.BoolVar( &configs.here, "here", false, helpHere );
	flag.BoolVar( &configs.noGit, "no-git", false, helpHere );
	flag.BoolVar( &configs.noTemplates, "ignore-templates", false, helpHere );
	flag.StringVar( &configs.license, "license", "", helpLicense );

	flag.Parse();
}

func ( configs *initConfigs ) parseConfigs( aProjectLanguage string ) bool {
	logs.DebugPrint( "Parse Configs" );

	lProjectLanguage := strings.ToLower( aProjectLanguage );
	lGlobalConfigFile := config.ConfigDir() + "config.toml";
	lLangKey := fmt.Sprintf( "init.%s", strings.TrimPrefix( lProjectLanguage, "." ));

	lGlobalConfigData, err := os.ReadFile( lGlobalConfigFile );
	if err != nil {
		logs.ErrorPrint( "Error reading config file:" + err.Error() );
		return false;
	}

	lTree, err := toml.Load( string(lGlobalConfigData) );
	if err != nil {
		fmt.Println( "Error loading TOML:", err );
		return false;
	}

	lSection := lTree.Get( lLangKey );
	if lSection == nil {
		fmt.Println( "Config not found for " + lProjectLanguage + " language" );
		return false;
	}

	lSectionMap := lSection.( *toml.Tree );

	// Parse command
	configs.command = "";
	if lParsedCommand := lSectionMap.Get( "command" ); lParsedCommand != nil {
		configs.command = lParsedCommand.( string );
	}

	// Parse license
	if configs.license == "" { // i.e. it has not been overridden
		if lParsedLicense := lSectionMap.Get( "license" ); lParsedLicense != nil {
			configs.license = lParsedLicense.( string );
		}
	}

	// Parse git_init
	if lParsedGitInit := lSectionMap.Get( "git_init" ); lParsedGitInit != nil {
		configs.noGit = lParsedGitInit.( bool );
	}

	// Parse git_ignore
	if lParsedGitIgnore := lSectionMap.Get( "git_ignore" ); lParsedGitIgnore != nil {
		if gitIgnores, ok := lParsedGitIgnore.( []any ); ok {
			for _, item := range gitIgnores {
				configs.gitIgnores = append( configs.gitIgnores, item.(string) );
			}
		}
	}

	// Parse templates
	if lParsedTemplates, ok := lSectionMap.Get( "templates" ).( []*toml.Tree ); ok {
		for _, tree := range lParsedTemplates {
			if title, titleOk := tree.Get( "title" ).( string ); titleOk {
				if destination, destOk := tree.Get( "destination" ).( string ); destOk {
					configs.templates = append(
						configs.templates,
						templates{ title: title, destination: destination },
					);
				}
			}
		}
	}

	// Parse directories
	if lParsedDirectories := lSectionMap.Get( "directories" ); lParsedDirectories != nil {
		if directories, ok := lParsedDirectories.( []any ); ok {
			for _, dir := range directories {
				configs.directories = append( configs.directories, dir.(string) );
			}
		}
	}

	return true;
}


//- Public Calls -----------------------------------------------------------------------------------

func Init() {
	var lConfigs initConfigs;

	if len( os.Args ) < 3 {
		logs.ErrorPrint( "Missing required arguments" );
		fmt.Println("");
		logs.HelpCommand( "init" );

		return;
    }

	lOriginalArgs := os.Args;
	os.Args = append( []string{ lOriginalArgs[0] }, lOriginalArgs[3:]... );

	// lOriginalArgs[1] = project name
	// lOriginalArgs[2] = project language

	lConfigs.parseInput();
	lConfigs.parseConfigs( lOriginalArgs[2] );

	// run commands
	if !lConfigs.here {
		// mkdir lOriginalArgs[1];
		// cd lOriginalArgs[1]
	}

	if !lConfigs.noTemplates {
		// for template in lConfigs.templates cp
	}

	if !lConfigs.noGit {
		// git init
		// cp CONFIG_DIR/ignores/language ./.gitginore
		// git commit init: project
	}
}

