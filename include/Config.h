
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
    char       *Title;
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
    char      **Extensions;
    char       *Build;
    char       *Run;
    char       *Clean;
    char       *Debug;
    char       *Test;
    bool        RootCD;
} ConfigDev;

typedef struct {
    ConfigLocal Local;
    ConfigRepo  Repository;
    ConfigInit *Init;
    ConfigDev  *Dev;
} Configs;


//- Declarations -----------------------------------------------------------------------------------

Configs *configs_parse( ConfigDomain );
void configs_free( Configs * );

ConfigDev *configs_dev( Configs * );

