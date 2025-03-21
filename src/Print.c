
#include "Debug.h"
#include "Print.h"

//- Helper Functions -------------------------------------------------------------------------------

void print_description( const char *aFlag, const char *aDescription )
{
    printf( "%s    --%-31s %s%s%s\n",
            TERM_GREEN, aFlag, TERM_GRAY, aDescription, TERM_RESET );
}

void print_command( const char *aCmd, const char *aFlags, const char *aDescription )
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

    print_usage( aSpecific );
}

void print_usage( Commands aSpecific )
{
    switch ( aSpecific )
    {
        case NONE:
            break;

        case RUN:
            print_command( "run", "{Args}", "Run programme with optional args" );
            printf( "\nFLAGS:\n" );
            print_description( "global", "Override local config with global settings" );
            break;

        case COMPILE:
            print_command( "compile", "{Flags}", "Compile programme with optional flags" );
            printf( "\nFLAGS:\n" );
            print_description( "global", "Override local config with global settings" );
            break;

        case BUN:
            print_command( "bun", "{Args}", "Build and run programme" );
            printf( "\nFLAGS:\n" );
            print_description( "flags [Flags]", "Specify build flags" );
            print_description( "args [Args]", "Specify run args" );
            print_description( "global", "Use global config, overriding local" );
            print_description( "global", "Override local config with global settings" );
            break;

        case INIT:
            print_command( "init", "[ProjectName] [Language]", "Initialise a new project." );
            printf( "\nFLAGS:\n" );
            print_description( "here", "Create project in current directory" );
            print_description( "licence", "Specify licence to use" );
            print_description( "no-git", "Omit Git initialisation" );
            print_description( "git-ignore [Args]", "Add to Git ignore list" );
            print_description( "git-ignore-only [Args]", "Create Git ignore list with specified args" );
            print_description( "template [Title] [Destination]", "Add template code to project" );
            print_description( "ignore-template", "Omit template code" );
            break;

        case TEMPLATE:
            print_command( "template", "{Args}", "Manage template files for projects." );
            printf( "\nFLAGS:\n" );
            print_description( "list", "List available templates" );
            print_description( "print-dir", "Show template directory" );
            print_description( "add [Title] [Destination]", "Add template code to project" );
            print_description( "new [Title]", "Create new template file" );
            break;

        case CLEAN:
            print_command( "clean", "", "Clean build files" );
            printf( "\nFLAGS:\n" );
            print_description( "global", "Override local config with global settings" );
            break;

        case TEST:
            print_command( "test", "", "Run test files" );
            printf( "\nFLAGS:\n" );
            print_description( "global", "Override local config with global settings" );
            break;

        case DEBUG:
            print_command( "debug", "", "Run project debugger" );
            printf( "\nFLAGS:\n" );
            print_description( "global", "Override local config with global settings" );
            break;

        case LICENSE:
            print_command( "licence", "{Args}", "Manage licences for projects." );
            printf( "\nFLAGS:\n" );
            print_description( "list", "List available licences" );
            print_description( "print-dir", "Show licence directory" );
            print_description( "replace [Title]", "Replace current licence" );
            print_description( "new [Title]", "Create new licence file" );
            break;

        case CONFIG:
            print_command( "config", "", "Manage configuration files for the dev utility." );
            printf( "\nFLAGS:\n" );
            print_description( "local", "Edit or create local config file" );
            break;

        case HELP:
            print_command( "help", "", "Display help menu" );
            break;

        case VERSION:
            print_command( "version", "", "Display utility version" );
            break;
    }
}

