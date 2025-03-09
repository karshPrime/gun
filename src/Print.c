
#include "Debug.h"
#include "Print.h"

void print_usage( void )
{
    debug( "print usage" );
}

void print_help( void )
{
    debug( "print help" );
}

void print_version( void )
{
    printf( "Version: %s\n", VERSION );
}

