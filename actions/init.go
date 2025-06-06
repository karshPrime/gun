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
};

type initConfigs struct {
	here		bool;
	command		string;
	license		string;
	noGit		bool;
	gitMessage	string;
	gitIgnores	[]string;
	noTemplates bool;
	templates	[]template;
	directories	[]string;
	files		[]string;
	copyConfig	bool;
	localConfig string;
};


//- Private Helpers --------------------------------------------------------------------------------

func ( configs *initConfigs ) parseInput() {
	flag.BoolVar( &configs.here, "here", false, helpHere );
	flag.BoolVar( &configs.noGit, "no-git", false, helpHere );
	flag.BoolVar( &configs.noTemplates, "no-templates", false, helpHere );
	flag.StringVar( &configs.license, "license", "", helpLicense );

	flag.Usage = func() {
		logs.HelpCommand( "init", false );
	};

	flag.Parse();
}

func ( configs *initConfigs ) parseConfigs( aProjectLanguage string ) bool {
	logs.DebugPrint( "Parse Configs" );

	// define initial list
	configs.gitIgnores = []string{ "# .gitignore", "" };
	configs.files = []string{ "README.md" }; // always create readme
	configs.command = "";

	// open the config.toml file
	lGlobalConfigFile := config.ConfigDir() + "config.toml";
	lGlobalConfigData, err := os.ReadFile( lGlobalConfigFile );
	if err != nil {
		logs.ErrorPrint( "Unable to read config file:", err );
		return false;
	}

	// parse config.toml file as toml
	lTree, err := toml.Load( string(lGlobalConfigData) );
	if err != nil {
		logs.ErrorPrint( "Unable to load TOML:", err );
		return false;
	}

	lReadCopyActions := func() {
		lSection := lTree.Get( "local" );
		if lSection == nil {
			return;
		}

		lSectionMap := lSection.( *toml.Tree );

		// Copy Config
		configs.copyConfig = false;
		if lParsedCommand := lSectionMap.Get( "copy_config" ); lParsedCommand != nil {
			configs.copyConfig = lParsedCommand.( bool );

			if lParsedCommand := lSectionMap.Get( "config_title" ); lParsedCommand != nil {
				configs.localConfig = lParsedCommand.( string );
			} else  {
				configs.localConfig = "commands";
			}
		}
	}

	// read data from the file: get to choose what toml block to read
	lReadGlobalConfig := func( aLangKey string ) int {
		lSection := lTree.Get( aLangKey );
		if lSection == nil {
			return 1;
		}

		lSectionMap := lSection.( *toml.Tree );

		// Parse command
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

			// Parse git init commit message
			if lParsedGitMsg := lSectionMap.Get( "git_message" ); lParsedGitMsg != nil {
				configs.gitMessage = lParsedGitMsg.( string );
			} else {
				configs.gitMessage = "init: project";
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

		// Parse file copy action
		lReadCopyActions();

		// Parse directories
		if lParsedDirectories := lSectionMap.Get( "directories" ); lParsedDirectories != nil {
			if directories, ok := lParsedDirectories.( []any ); ok {
				for _, dir := range directories {
					configs.directories = append( configs.directories, dir.(string) );
				}
			}
		}

		// Parse files
		if lParsedFiles := lSectionMap.Get( "files" ); lParsedFiles != nil {
			if files, ok := lParsedFiles.( []any ); ok {
				for _, dir := range files {
					configs.files = append( configs.files , dir.(string) );
				}
			}
		}
		return 0;
	}

	// read the default [init] defined config, and set those as values
	lReadGlobalConfig( "init" ); // ignore error messages

	// override [init] defines with specific defined
	lProjectLanguage := strings.ToLower( aProjectLanguage );
	lLangKey := fmt.Sprintf( "init.%s", strings.TrimPrefix( lProjectLanguage, "." ));
	switch ( lReadGlobalConfig( lLangKey )) {
		case 1:
			logs.ErrorPrint( "Config not found for " + lProjectLanguage + " language" );
			return false;

		default:
			return true;
	}
}

func replaceConfigPlaceholders( aString *string, aProjectName string, aProjectLang string ) {
	lScriptPath := config.ConfigDir() + "scripts/";

	logs.DebugPrint(  *aString,  aProjectName, aProjectLang );

	*aString = strings.Replace( *aString, "%PROJECT_NAME%", aProjectName, -1 );
	*aString = strings.Replace( *aString, "%PROJECT_LANGUAGE%", aProjectLang, -1 );
	*aString = strings.Replace( *aString, "%SCRIPT% ", lScriptPath, -1 );
}


//- Public Calls -----------------------------------------------------------------------------------

func Init() {
	var lConfigs initConfigs;

	if len( os.Args ) < 3 {
		logs.ErrorPrint( "Missing required arguments" );
		fmt.Println(".");
		logs.HelpCommand( "init", false );

		return;
    }

	lOriginalArgs := os.Args;
	os.Args = append( []string{ lOriginalArgs[0] }, lOriginalArgs[3:]... );

	// lOriginalArgs[1] = project name
	// lOriginalArgs[2] = project language


	lConfigs.parseInput();
	lConfigs.parseConfigs( lOriginalArgs[2] );
	replaceConfigPlaceholders( &lConfigs.command, lOriginalArgs[1], lOriginalArgs[2]);

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
				logs.ErrorPrint( "Project directory already exists" );
				fmt.Println( "." );
				return;
			} else {
				logs.ErrorPrint( "Unable to create directory: %v", err );
				fmt.Println( "." );
				return;
			}
		}

		err = os.Chdir( lOriginalArgs[1] );
		if err != nil {
			logs.ErrorPrint( "Unable to change directory: %v", err );
			fmt.Println( "." );
			return;
		}
	}

	// Copy dev commands into command
	if lConfigs.copyConfig {
		lAllCommands := []Triggers{ RUN, BUILD, TEST, DEBUG, CLEAN };
		var lTConfig triggerConfigs;
		lTConfig.global = true;
		var lOutputText []string;

		for _, cmd := range lAllCommands {
			lTConfig.parseConfigs( cmd, lOriginalArgs[2] );

			if lTConfig.command != "" {
				lOutputText = append(lOutputText, fmt.Sprintf(
					"\t%-7s = '%s',", triggersKey(cmd), strings.TrimSpace(lTConfig.command),
				));

				lTConfig.command = "";
			}
		}

		lOutputText = append( lOutputText, fmt.Sprintf( "\tcd_root = %t,", lTConfig.cdRoot ) );

		lFileContent := "\n-- " + config.GunURL + "\n\nreturn {\n" +
			strings.Join( lOutputText, "\n" ) + "\n}\n\n";

		err := os.WriteFile( lConfigs.localConfig, []byte(lFileContent), 0644 )
		if err != nil {
			logs.ErrorPrint( "Unable to copy dev commands to local file", err );
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
		replaceConfigPlaceholders( &lFileName, lOriginalArgs[1], lOriginalArgs[2] );
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
		fmt.Fprintln( os.Stderr, lResult );
	}

	// Create git repo
	if !lConfigs.noGit {
		lResult, err := SysRun( "git init" );
		if err {
			logs.ErrorPrint( err );
			goto FinishInit;
		}
		fmt.Fprintln( os.Stderr, lResult );

		// .gitignore
		file, lFileErr := os.OpenFile( ".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644 )
		if lFileErr != nil {
			logs.ErrorPrint( "Unable to open or create .gitignore: %v", err );
			goto FinishInit;
		}
		defer file.Close();
		for _, line := range lConfigs.gitIgnores {
			replaceConfigPlaceholders( &line, lOriginalArgs[1], lOriginalArgs[2] );
			if _, err := file.WriteString( line + "\n" ); err != nil {
				goto FinishInit;
			}
		}

		lResult, err = SysRun( "git add -A" );
		if err {
			logs.ErrorPrint( err );
			goto FinishInit;
		}

		replaceConfigPlaceholders( &lConfigs.gitMessage, lOriginalArgs[1], lOriginalArgs[2] );
		lResult, err = SysRun( "git commit -m \"" + lConfigs.gitMessage + "\"" );
		if err {
			logs.DebugPrint( "git commit -m \"" + lConfigs.gitMessage + "\"" );
			logs.ErrorPrint( err );
			goto FinishInit;
		}
		fmt.Fprintln( os.Stderr, lResult );
	}

	FinishInit:
	fmt.Println( "./" + strings.ReplaceAll(lOriginalArgs[1], " ", "\\ " ) );
}

