
/* Config.h
 *
 *
 */

#pragma once

//- Environment Variables --------------------------------------------------------------------------

#if defined(_WIN32) && defined(_WIN64)
    #define CONFIG_DIR ""

#else // macOS or Linux
    #define CONFIG_DIR "~/.config/devconfig"

#endif


//- Custom Datatypes -------------------------------------------------------------------------------

typedef enum { false, true } bool;
typedef enum { LOCAL, GLOBAL } ConfigDomain;
typedef unsigned int uint;

typedef struct {
    bool        CopyConfigLocal;
    bool        GitInit;
    char       *Command;
    char       *License;
    char      **GitIgnore;
    char      **Directories;
    char      **Files;
    char      **Templates;
    struct {
        char   *Username;
        char   *Hostname;
        bool    SSH;
        bool    Commit;
    } ConfigRepo;
} ConfigInit;

typedef struct {
    char       *Build;
    char       *Run;
    char       *Clean;
    char       *Debug;
    char       *Test;
    bool        RootCD;
} ConfigDev;


//- Declarations -----------------------------------------------------------------------------------

ConfigInit *config_parse_init( char *, bool );
char **config_parse_template( void );
ConfigDev *config_parse_dev( bool );

void config_free_init( ConfigInit * );
void config_free_template( char ** );
void config_free_dev( ConfigDev * );

