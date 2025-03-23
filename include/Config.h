
/* Config.h
 *
 *
 */

#pragma once

//- Custom Datatypes -------------------------------------------------------------------------------

typedef enum { false, true } bool;
typedef enum { LOCAL, GLOBAL } ConfigDomain;
typedef unsigned int uint;

typedef struct {
    bool        Copy;
    char       *Title;
} ConfigLocal;

typedef struct {
    char       *Username;
    char       *Hostname;
    bool        SSH;
    bool        Commit;
} ConfigRepo;

typedef struct {
    bool        GitInit;
    char      **GitIgnore;
    char      **Directories;
    char      **Files;
    char       *Command;
    char       *License;
    struct {
        char   *Title;
        char   *Destination;
    } *Templates;
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

