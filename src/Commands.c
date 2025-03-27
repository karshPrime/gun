
#include "Commands.h"
#include "Debug.h"

void cmd_init( bool aLocalConfig, InitArgs aArgs )
{
    debug( "cmd_init() called" );

    ConfigInit *lConfigs = config_parse_init();

    //

    config_free_init( lConfigs );
}

void cmd_license( RecordsArgs aArgs, char *aReplaceTitle )
{
    debug( "cmd_license() called" );

    // get all files in set directory
}

void cmd_template( RecordsArgs aArgs, char **aTemplates, bool aManage )
{
    debug( "cmd_template() called" );

    char **lTemplates = config_parse_template();

    //

    config_free_template( lTemplates );
}

