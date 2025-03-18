
#include "Parse.h"
#include <string.h>

//--------------------------------------------------------------------------------------------------

bool parse_check_value( const char *aArg, const char *aFull, const char aShort )
{
    return strcmp( aArg, aFull ) == 0 || (*aArg == aShort && aArg[1] == '\0');
}

uint parse_input_count( uint aIndex, const uint aLen, char * aArg[] )
{
    uint Result = 0;

    for ( uint i = aIndex+1; i < aLen; i++ )
    {
        if ( aArg[i][0] == '-' )
            break;

        Result++;
    }

    return Result + aIndex + 1;
}

