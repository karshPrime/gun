
#pragma once

#include "Config.h"

//- Custom Datatypes -------------------------------------------------------------------------------

typedef struct {
    char   *Title;
    char   *Destination;
} Move;

typedef struct {
    bool    Here;
    bool    NoGit;
    char   *License;
    char   *GitIgnore;
    char   *GitOnlyIgnore;
    Move   *Templates;
    bool    TemplateIgnore;
} InitArgs;

typedef struct {
    bool    Lists;
    bool    PrintDir;
    char   *Title;
} RecordsArgs;


//- Declarations -----------------------------------------------------------------------------------

void cmd_init( Configs *, InitArgs );
void cmd_license( Configs *, RecordsArgs, char * );
void cmd_template( Configs *, RecordsArgs, Move * );

