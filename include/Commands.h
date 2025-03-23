
#pragma once

#include "Config.h"

//- Custom Datatypes -------------------------------------------------------------------------------

typedef struct {
    bool    Here;
    bool    NoGit;
    char   *License;
    char  **GitIgnore;
    char  **GitOnlyIgnore;
    char  **Templates;
    bool    TemplateIgnore;
} InitArgs;

typedef struct {
    bool    Lists;
    bool    PrintDir;
    char   *New;
} RecordsArgs;


//- Declarations -----------------------------------------------------------------------------------

void cmd_init( Configs *, InitArgs );
void cmd_license( Configs *, RecordsArgs, char * );
void cmd_template( Configs *, RecordsArgs, char **, bool );

