
// main.c

#include "Config.h"
#include "Debug.h"
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

    return 0;
}

