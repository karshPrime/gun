
/* Actions.c
 *
 *
 */

#include "Actions.h"
#include "Debug.h"

#include <stdlib.h>
#include <unistd.h>

//- Helper Functions -------------------------------------------------------------------------------

void _run_cmd( const bool aRootCD, const char *aCommand )
{
    if ( aRootCD )
    {
        if ( !system( "git rev-parse --is-inside-work-tree > /dev/null 2>&1" ) )
        {
            FILE *lTermOutput = popen( "git rev-parse --show-toplevel", "r" );
            if ( lTermOutput == NULL )
            {
                log_error( "popen" );
                return;
            }

            char lTermBuffer[256];
            if ( fgets( lTermBuffer, 256, lTermOutput ) )
            {
                lTermBuffer[strcspn( lTermBuffer, "\n" )] = 0;
                if ( chdir( lTermBuffer ) )
                {
                    log_error( "chdir" );
                    return;
                }
            }

            pclose( lTermOutput );
        }
    }

    if ( aCommand )
    {
        debug( "%s", aCommand );
        system( aCommand );
    }
    else
    {
        log_error( "Command not set in config" );
    }
}


//- External Calls ---------------------------------------------------------------------------------

void action_bun( const bool aLocalConfig, char *aArgs[], char *aFlags[] )
{
    debug( "action_bun() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    _run_cmd( lConfigs->RootCD, lConfigs->Build );
    _run_cmd( false, lConfigs->Run ); // dont cd again; bloat steps

    config_free_dev( lConfigs );
}

void action_run( const bool aLocalConfig, char *aArgs[] )
{
    debug( "action_run() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    _run_cmd( lConfigs->RootCD, lConfigs->Run );

    config_free_dev( lConfigs );
}

void action_compile( const bool aLocalConfig, char *aFlags[] )
{
    debug( "action_compile() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    _run_cmd( lConfigs->RootCD, lConfigs->Build );

    config_free_dev( lConfigs );
}

void action_clean( const bool aLocalConfig )
{
    debug( "action_clean() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    _run_cmd( lConfigs->RootCD, lConfigs->Clean );

    config_free_dev( lConfigs );
}

void action_debug( const bool aLocalConfig )
{
    debug( "action_debug() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    _run_cmd( lConfigs->RootCD, lConfigs->Debug );

    config_free_dev( lConfigs );
}

void action_test( const bool aLocalConfig )
{
    debug( "action_test() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    _run_cmd( lConfigs->RootCD, lConfigs->Test );

    config_free_dev( lConfigs );
}

void action_config( const bool aLocal )
{
    debug( "action_config() called" );

    char lCommand[256];

    const char *lConfig = aLocal ? "./" : CONFIG_DIR;
    sprintf( lCommand, "$EDITOR %s/config.toml", lConfig );

    system( lCommand );
}

