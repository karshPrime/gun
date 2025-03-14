
// main.c

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

    for ( uint i = 0; i < argc; i++ )
    {
        if      ( parse_check_value( argv[i], "help"   , "h" ) ) { print_help();     }
        else if ( parse_check_value( argv[i], "version", "v" ) ) { print_version();  }
    }


    return 0;
}

