
#include "Debug.h"
#include "Print.h"

//- Helper Functions -------------------------------------------------------------------------------

void print_description( const char *aFlag, const char *aDescription )
{
}

void print_command( const char *aCmd, const char *aFlags, const char *aDescription )
{
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
            break;

        case COMPILE:
            break;

        case BUN:
            break;

        case INIT:
            break;

        case TEMPLATE:
            break;

        case CLEAN:
            break;

        case TEST:
            break;

        case DEBUG:
            break;

        case LICENSE:
            break;

        case CONFIG:
            break;

        case HELP:
            break;

        case VERSION:
            break;
    }
}

