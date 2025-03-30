
/* Config.c
 *
 *
 */

#include "Config.h"
#include "Debug.h"
#include "tomlc99/toml.h"

#include <stdlib.h>

//- Helper Functions -------------------------------------------------------------------------------



//- External Calls ---------------------------------------------------------------------------------

ConfigInit *config_parse_init( char *aLanguage, bool aCopyTemplate )
{
    debug( "Reading configs for init" );

    ConfigInit *Result;
    Result->Templates = config_parse_template();

    return Result;
}

char **config_parse_template( void )
{
    debug( "Reading configs for tempaltes" );

    char **Result;

    return Result;
}

ConfigDev *config_parse_dev( bool aIsLocal )
{
    debug( "Reading configs for dev" );

    ConfigDev *Result;

    return Result;
}

void config_free_init( ConfigInit *aInitObj )
{
    debug( "freeing allocated init" );

    config_free_template( aInitObj->Templates );
    free( aInitObj->GitIgnore );
    free( aInitObj->Directories );
    free( aInitObj->Files );
    free( aInitObj );
}

void config_free_template( char **aTemplateObj )
{
    debug( "freeing allocated templates" );

    free( aTemplateObj );
}

void config_free_dev( ConfigDev *aDevObj )
{
    debug( "freeing allocated dev" );

    free( aDevObj );
}

