
#include "Config.h"
#include "Debug.h"
#include <stdlib.h>

Configs *configs_parse( ConfigDomain aDomain )
{
    debug( "*configs_parse() called" );
    Configs *Result = malloc( sizeof( Configs ) );

    return Result;
}

void configs_free( Configs *aConfigs )
{
    free( aConfigs );
    debug( "configs_free() called" );
}

ConfigDev *configs_dev( Configs *aConfigs )
{
    debug( "*configs_dev() called" );
    return aConfigs->Dev;
}

