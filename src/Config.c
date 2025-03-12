
#include "Config.h"
#include <stdlib.h>

Configs *configs_parse( void )
{
    Configs *Result = malloc( sizeof( Configs ) );

    return Result;
}


void configs_free( Configs *aConfigs )
{
    free( aConfigs );
}

ConfigDev *configs_dev( Configs *aConfigs )
{
    return aConfigs->Dev;
}

