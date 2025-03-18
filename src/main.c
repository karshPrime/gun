
// main.c

#include "Actions.h"
#include "Commands.h"
#include "Config.h"
#include "Debug.h"
#include "Parse.h"
#include "Print.h"

const char *VERSION = "0.0.1";

int main( int argc, char *argv[] )
{
    if ( argc < 2 )
    {
        log_error( "Missing Arguments" );
        print_usage();
        return 1;
    }

    const char *COMMAND = argv[1];
    Configs *lConfigs = configs_parse();

    if ( parse_check_value( COMMAND, "bun", 'b' ) )
    {
        debug( "called bun" );
        char *lFlags[1]; // size 1 is temporal
        char *lArgs[1];

        action_bun( lArgs, lFlags );
    }

    else if ( parse_check_value( COMMAND, "run", 'r' ) )
    {
        char *lArgs[1];

        action_run( lArgs );
    }

    else if ( parse_check_value( COMMAND, "compile", 'c' ) )
    {
        char *lFlags[1];

        action_run( lFlags );
    }

    else if ( parse_check_value( COMMAND, "debug",  'd' ) ) { action_debug(); }
    else if ( parse_check_value( COMMAND, "test",   't' ) ) { action_test();  }
    else if ( parse_check_value( COMMAND, "clean",  'x' ) ) { action_clean(); }

    else if ( parse_check_value( COMMAND, "init", 'i' ) )
    {
        InitArgs lInitArgs = {
            .Here = false,
            .NoGit = false,
            .TemplateIgnore = false
        };

        // parse remaining flags

        cmd_init( lConfigs, lInitArgs );
    }

    else if ( parse_check_value( COMMAND, "license", 'l' ) )
    {
        RecordsArgs lRecords;
        char *lName;

        cmd_license( lConfigs, lRecords, lName );
    }

    else if ( parse_check_value( COMMAND, "tempalte", 'T' ) )
    {
        RecordsArgs lRecords;
        Move *lMove;

        cmd_template( lConfigs, lRecords, lMove );
    }

    else if ( parse_check_value( COMMAND, "config", 'z' ) )
    {
        bool lIsLocal = false;

        action_config( lIsLocal );
    }

    else if ( parse_check_value( COMMAND, "help",    "h" ) ) { print_help();    }
    else if ( parse_check_value( COMMAND, "version", "v" ) ) { print_version(); }

    else
    {
        log_error( "Invalid Argument." );
        print_usage();
    }

    return 0;
}

