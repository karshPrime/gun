
/* Print.c
 *
 *
 */

#include "Print.h"
#include "Debug.h"

//- Helper Functions -------------------------------------------------------------------------------

void _print_description( const char *aFlag, const char *aDescription )
{
    printf( "%s    --%-22s %s%s%s\n",
            TERM_GREEN, aFlag, TERM_GRAY, aDescription, TERM_RESET );
}

void _print_command( const char *aCmd, const char *aFlags, const char *aDescription )
{
    printf( "%s\n\nUSAGE:\n   %s$ %sdev %s%s %s%s%s\n",
            aDescription, TERM_BLUE, TERM_RED, TERM_YELLOW, aCmd, TERM_BLUE, aFlags, TERM_RESET );
}


//- External Calls ---------------------------------------------------------------------------------

void print_version( void )
{
    printf( "Version: %s\n", CURRENT_VERSION );
}

void print_help( Commands aSpecific )
{
    debug( "print help" );

    switch ( aSpecific )
    {
        case NONE:
            break;

        case RUN:
            _print_command( "run", "{Args}", "Run programme with optional args" );
            printf( "\nFLAGS:\n" );
            _print_description( "global", "Override local config with global settings" );
            break;

        case COMPILE:
            _print_command( "compile", "{Flags}", "Compile programme with optional flags" );
            printf( "\nFLAGS:\n" );
            _print_description( "global", "Override local config with global settings" );
            break;

        case BUN:
            _print_command( "bun", "", "Build and run programme" );
            printf( "\nFLAGS:\n" );
            _print_description( "flags [Flags]", "Specify build flags" );
            _print_description( "args [Args]", "Specify run args" );
            _print_description( "global", "Override local config with global settings" );
            break;

        case INIT:
            _print_command( "init", "[ProjectName] [Language]", "Initialise a new project." );
            printf( "\nFLAGS:\n" );
            _print_description( "here", "Create project in current directory" );
            _print_description( "licence [Title]", "Specify licence to use" );
            _print_description( "no-git", "Omit Git initialisation" );
            _print_description( "git-ignore [Args]", "Add to Git ignore list" );
            _print_description( "git-ignore-only [Args]",
                    "Create ignore list with only the specified args" );
            _print_description( "template [Templates]", "Add template code to project" );
            _print_description( "ignore-template", "Omit template code" );
            break;

        case TEMPLATE:
            _print_command( "template", "{Args}", "Manage template files for projects." );
            printf( "\nFLAGS:\n" );
            _print_description( "list", "List available templates" );
            _print_description( "add [Templates]", "Add template code to project" );
            _print_description( "new [Title]", "Create new template file" );
            _print_description( "manage", "Manage the record of all saved templates" );
            _print_description( "print-dir", "Show template directory" );
            break;

        case CLEAN:
            _print_command( "clean", "", "Clean build files" );
            printf( "\nFLAGS:\n" );
            _print_description( "global", "Override local config with global settings" );
            break;

        case TEST:
            _print_command( "test", "", "Run test files" );
            printf( "\nFLAGS:\n" );
            _print_description( "global", "Override local config with global settings" );
            break;

        case DEBUG:
            _print_command( "debug", "", "Run project debugger" );
            printf( "\nFLAGS:\n" );
            _print_description( "global", "Override local config with global settings" );
            break;

        case LICENSE:
            _print_command( "licence", "{Args}", "Manage licences for projects." );
            printf( "\nFLAGS:\n" );
            _print_description( "list", "List available licences" );
            _print_description( "replace [Title]", "Replace current licence" );
            _print_description( "new [Title]", "Create new licence file" );
            _print_description( "print-dir", "Show licence directory" );
            break;

        case CONFIG:
            _print_command( "config", "", "Manage configuration files for the dev utility." );
            printf( "\nFLAGS:\n" );
            _print_description( "local", "Edit or create local config file" );
            break;

        case HELP:
            _print_command( "help", "", "Display help menu" );
            break;

        case VERSION:
            _print_command( "version", "", "Display utility version" );
            break;
    }
}

