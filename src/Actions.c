
/* Actions.c
 *
 *
 */

#include "Actions.h"
#include "Debug.h"

//- Helper Functions -------------------------------------------------------------------------------



//- External Calls ---------------------------------------------------------------------------------

void action_bun( bool aLocalConfig, char *aArgs[], char *aFlags[] )
{
    debug( "action_bun() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    //

    config_free_dev( lConfigs );
}

void action_run( bool aLocalConfig, char *aArgs[] )
{
    debug( "action_run() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    //

    config_free_dev( lConfigs );
}

void action_compile( bool aLocalConfig, char *aFlags[] )
{
    debug( "action_compile() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    //

    config_free_dev( lConfigs );
}

void action_clean( bool aLocalConfig )
{
    debug( "action_clean() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    //

    config_free_dev( lConfigs );
}

void action_debug( bool aLocalConfig )
{
    debug( "action_debug() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    //

    config_free_dev( lConfigs );
}

void action_test( bool aLocalConfig )
{
    debug( "action_test() called" );

    ConfigDev *lConfigs = config_parse_dev( aLocalConfig );

    //

    config_free_dev( lConfigs );
}

void action_config( bool aLocal )
{
    debug( "action_config() called" );
}

