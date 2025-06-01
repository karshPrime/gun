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

type template struct {
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
	templates	[]template;
	directories	[]string;
	files		[]string;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *initConfigs ) parseInput() {
	flag.BoolVar( &configs.here, "here", false, helpHere );
	flag.BoolVar( &configs.noGit, "no-git", false, helpHere );
	flag.BoolVar( &configs.noTemplates, "no-templates", false, helpHere );
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
		logs.ErrorPrint( "Unable to read config file:", err );
		return false;
	}

	lTree, err := toml.Load( string(lGlobalConfigData) );
	if err != nil {
		logs.ErrorPrint( "Unable to load TOML:", err );
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

	if !configs.noGit {
		// Parse git_init
		if lParsedGitInit := lSectionMap.Get( "git_init" ); lParsedGitInit != nil {
			configs.noGit = !lParsedGitInit.( bool );
		}

		// Parse git_ignore
		if lParsedGitIgnore := lSectionMap.Get( "git_ignore" ); lParsedGitIgnore != nil {
			if gitIgnores, ok := lParsedGitIgnore.( []any ); ok {
				for _, item := range gitIgnores {
					configs.gitIgnores = append( configs.gitIgnores, item.(string) );
				}
			}
		}
	}

	// Parse templates
	if !configs.noTemplates {
		if lParsedTemplates, ok := lSectionMap.Get( "templates" ).( []*toml.Tree ); ok {
			for _, tree := range lParsedTemplates {
				if title, titleOk := tree.Get( "title" ).( string ); titleOk {
					if destination, destOk := tree.Get( "destination" ).( string ); destOk {
						configs.templates = append(
							configs.templates,
							template{ title: title, destination: destination },
						);
					}
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

	// Parse files
	configs.files = []string{ "README.md" }; // always create readme
	if lParsedFiles := lSectionMap.Get( "files" ); lParsedFiles != nil {
		if files, ok := lParsedFiles.( []any ); ok {
			for _, dir := range files {
				configs.files = append( configs.files , dir.(string) );
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

	// logs.ErrorPrint( "command\t\t:", lConfigs.command )
	// logs.ErrorPrint( "here\t\t:", lConfigs.here )
	// logs.ErrorPrint( "license\t\t:", lConfigs.license )
	// logs.ErrorPrint( "noGit\t\t:", lConfigs.noGit )
	// logs.ErrorPrint( "gitIgnores\t:", lConfigs.gitIgnores )
	// logs.ErrorPrint( "noTemplates\t:", lConfigs.noTemplates )
	// logs.ErrorPrint( "templates\t:", lConfigs.templates )
	// logs.ErrorPrint( "directories\t:", lConfigs.directories )

	// Create project directory and cd into it
	if !lConfigs.here {
		err := os.Mkdir( lOriginalArgs[1], 0755 );
		if err != nil {
			if os.IsExist(err) {
				logs.ErrorPrint("Project directory already exists" );
				return;
			} else {
				logs.ErrorPrint( "Unable to create directory: %v", err );
				return;
			}
		}

		err = os.Chdir( lOriginalArgs[1] );
		if err != nil {
			logs.ErrorPrint( "Unable to change directory: %v", err );
			return;
		}
	}

	// Create new directories
	for _, lDirName := range lConfigs.directories {
		err := os.Mkdir( lDirName, 0755 );
		if err != nil {
			if os.IsExist(err) {
				logs.WarningPrint( lDirName, "directory already exists. Skipping" );
			} else {
				logs.WarningPrint( "Unable to create directory %s: %v", lDirName, err );
			}
		}
	}

	// Create new files
	for _, lFileName := range lConfigs.files {
		file, err := os.Create( lFileName );
		if err != nil {
			if os.IsExist(err) {
				logs.WarningPrint( lFileName, "file already exists. Skipping" );
			} else {
				logs.WarningPrint( "Unable to create directory %s: %v", lFileName, err );
			}
		}
		file.Close();
	}

	// Copy all templates
	if !lConfigs.noTemplates {
		for _, lTemplate := range lConfigs.templates {
			lTitle := config.ConfigDir() + "templates/" + lTemplate.title;

			if Copy( lTitle, lTemplate.destination ) {
				logs.ErrorPrint( "Unable to copy template from %s to %s",
					lTemplate.title, lTemplate.destination );
			}
		}
	}

	// Copy License
	if lConfigs.license != "" {
		lLicensePath := config.ConfigDir() + "licenses/" + lConfigs.license;
		logs.DebugPrint( lLicensePath );
		Copy( lLicensePath, "./LICENSE" );
	}

	// Run init command
	lResult, lError := SysRun( lConfigs.command );
	if lResult == "" {
	} else if lError {
		logs.ErrorPrint( lResult );
	} else  {
		fmt.Println( lResult );
	}

	// Create git repo
	if !lConfigs.noGit {
		lResult, err := SysRun( "git init" );
		if err {
			logs.ErrorPrint( err );
			return;
		}
		fmt.Println( lResult );

		// write to .gitignore

		lResult, err = SysRun( "git add -A" );
		if err {
			logs.ErrorPrint( err );
			return;
		}

		lResult, err = SysRun( "git commit -m \"init: project\"" );
		if err {
			logs.ErrorPrint( err );
			return;
		}
		fmt.Println( lResult );
	}
}

