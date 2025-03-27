
// main.c

#include "Actions.h"
#include "Commands.h"
#include "Debug.h"
#include "Parse.h"
#include "Print.h"

#include <string.h>

const char *CURRENT_VERSION = "0.0.1";

int main( int argc, char *argv[] )
{
    if ( argc < 2 )
    {
        log_error( "Missing Arguments" );
        print_help( NONE );
        return 1;
    }

    const char *COMMAND = argv[1];
    bool lLocalConfig = true;

    if ( parse_check_value( COMMAND, "bun", 'b' ) )
    {
        int  lArgsIndex  = 0;
        int  lFlagsIndex = 0;
        bool lFlagsFound = false;
        char *lArgs[argc];
        char *lFlags[argc];

        for ( int i = 2; i < argc; i++ )
        {
            if ( strcmp( argv[i], "--args" ) == 0 )
            {
                lFlagsFound = false;
            }
            else if ( strcmp( argv[i], "--flags" ) == 0 )
            {
                lFlagsFound = true;
            }
            else if ( strcmp( argv[i], "--global" ) == 0 )
            {
                lLocalConfig = false;
            }
            else if ( lFlagsFound )
            {
                lFlags[lFlagsIndex++] = argv[i];
            }
            else
            {
                lArgs[lArgsIndex++] = argv[i];
            }
        }

        action_bun( lLocalConfig, lArgs, lFlags );
    }

    else if ( parse_check_value( COMMAND, "run", 'r' ) )
    {
        char *lArgs[argc];
        int lArgsIndex = 0;

        for ( int i = 2; i < argc; i++ )
        {
            if ( strcmp( argv[i], "--global" ) == 0 )
            {
                lLocalConfig = false;
            }
            else
            {
                lArgs[lArgsIndex++] = argv[i];
            }
        }

        action_run( lLocalConfig, lArgs );
    }

    else if ( parse_check_value( COMMAND, "compile", 'c' ) )
    {
        char *lFlags[argc];
        int lFlagsIndex = 0;

        for ( int i = 2; i < argc; i++ )
        {
            if ( strcmp( argv[i], "--global" ) == 0 )
            {
                lLocalConfig = false;
            }
            else
            {
                lFlags[lFlagsIndex++] = argv[i];
            }
        }

        action_compile( lLocalConfig, lFlags );
    }

    else if ( parse_check_value( COMMAND, "debug", 'd' ) )
    {
        lLocalConfig = 2 > argc || strcmp( argv[2], "--global" );

        action_debug( lLocalConfig );
    }

    else if ( parse_check_value( COMMAND, "test", 't' ) )
    {
        lLocalConfig = 2 > argc || strcmp( argv[2], "--global" );

        action_test( lLocalConfig );
    }

    else if ( parse_check_value( COMMAND, "clean", 'x' ) )
    {
        lLocalConfig = 2 > argc || strcmp( argv[2], "--global" );

        action_clean( lLocalConfig );
    }

    else if ( parse_check_value( COMMAND, "init", 'i' ) )
    {
        if ( argc < 4 )
        {
            log_error( "Missing Project description" );
            print_help( INIT );

            return 1;
        }

        InitArgs lInitArgs = {
            .ProjectName = argv[2],
            .Language = argv[3],
            .Here = false,
            .NoGit = false,
            .TemplateIgnore = false
        };

        for ( int i = 4; i < argc; i++ )
        {
            if ( strcmp( argv[i], "--here" ) == 0 )
            {
                lInitArgs.Here = true;
            }
            else if ( i+1 < argc && strcmp( argv[i], "--license" ) == 0 )
            {
                lInitArgs.License = argv[i+1];
            }
            else if ( strcmp( argv[i], "--no-git" ) == 0 )
            {
                lInitArgs.NoGit = true;
            }
            else if ( i+1 < argc && strcmp( argv[i], "--git-ignore" ) == 0 )
            {
                const int lStartIndex = i;
                char *lArgs[argc-lStartIndex];

                while ( argv[i][0] != '-' )
                {
                    lArgs[i-lStartIndex] = argv[i];
                    i++;
                }

                lInitArgs.GitIgnore = lArgs;
            }
            else if ( i+1 < argc && strcmp( argv[i], "--git-ignore-only" ) == 0 )
            {
                const int lStartIndex = i;
                char *lArgs[argc-lStartIndex];

                while ( argv[i][0] != '-' )
                {
                    lArgs[i-lStartIndex] = argv[i];
                    i++;
                }

                lInitArgs.GitOnlyIgnore = lArgs;
            }
            else if ( i+1 < argc && strcmp( argv[i], "--template" ) == 0 )
            {
                const int lStartIndex = i;
                char *lArgs[argc-lStartIndex];

                while ( argv[i][0] != '-' )
                {
                    lArgs[i-lStartIndex] = argv[i];
                    i++;
                }

                lInitArgs.Templates = lArgs;
            }
        }

        cmd_init( lLocalConfig, lInitArgs );
    }

    else if ( parse_check_value( COMMAND, "license", 'l' ) )
    {
        RecordsArgs lRecords;
        char *lName;

        if ( argc <=  2 )
        { }
        else if ( strcmp( argv[2], "--lists" ) == 0 )
        {
            lRecords.Lists = true;
        }
        else if ( 3 < argc && strcmp( argv[2], "--new" ) == 0 )
        {
            lRecords.New = argv[3];
        }
        else if ( 3 < argc && strcmp( argv[2], "--replace" ) == 0 )
        {
            lName = argv[3];
        }
        else if ( strcmp( argv[2], "--print-dir" ) == 0 )
        {
            lRecords.PrintDir = true;
        }

        cmd_license( lRecords, lName );
    }

    else if ( parse_check_value( COMMAND, "template", 'T' ) )
    {
        RecordsArgs lRecords;
        char *lTemplates[argc];
        bool lManage = false;

        if ( argc <=  2 )
        { }
        else if ( strcmp( argv[2], "--lists" ) == 0 )
        {
            lRecords.Lists = true;
        }
        else if ( 3 < argc && strcmp( argv[2], "--new" ) == 0 )
        {
            lRecords.New = argv[3];
        }
        else if ( 3 < argc && strcmp( argv[2], "--add" ) == 0 )
        {
            for ( int i = 3; i < argc; i++ )
            {
                lTemplates[i-3] = argv[i];
            }
        }
        else if ( strcmp( argv[2], "--manage" ) == 0 )
        {
            lManage = true;
        }
        else if ( strcmp( argv[2], "--print-dir" ) == 0 )
        {
            lRecords.PrintDir = true;
        }

        cmd_template( lRecords, lTemplates, lManage );
    }

    else if ( parse_check_value( COMMAND, "config", 'z' ) )
    {
        bool lIsLocal = ( 2 < argc && strcmp( argv[2], "--local" ) == 0 );

        action_config( lIsLocal );
    }

    else if ( parse_check_value( COMMAND, "help", 'h' ) )
    {
        if ( argc < 3 ) print_help( NONE );
        else if ( strcmp( argv[2], "run" ) == 0 ) print_help( RUN );
        else if ( strcmp( argv[2], "bun" ) == 0 ) print_help( BUN );
        else if ( strcmp( argv[2], "init" ) == 0 ) print_help( INIT );
        else if ( strcmp( argv[2], "help" ) == 0 ) print_help( HELP );
        else if ( strcmp( argv[2], "test" ) == 0 ) print_help( TEST );
        else if ( strcmp( argv[2], "clean" ) == 0 ) print_help( CLEAN );
        else if ( strcmp( argv[2], "debug" ) == 0 ) print_help( DEBUG );
        else if ( strcmp( argv[2], "config" ) == 0 ) print_help( CONFIG );
        else if ( strcmp( argv[2], "license" ) == 0 ) print_help( LICENSE );
        else if ( strcmp( argv[2], "version" ) == 0 ) print_help( VERSION );
        else if ( strcmp( argv[2], "compile" ) == 0 ) print_help( COMPILE );
        else if ( strcmp( argv[2], "template" ) == 0 ) print_help( TEMPLATE );
        else
        {
            log_error( "Invalid Option %s", argv[2] );
            print_help( NONE );
        }
    }

    else if ( parse_check_value( COMMAND, "version", 'v' ) )
    {
        print_version();
    }

    else
    {
        log_error( "Invalid Argument." );
        print_help( NONE );
    }
}

