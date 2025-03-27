
/* Actions.h
 * Actions supported by the utility, such as to compile/run/bun/clean etc.
 * This code takes the user arguments and runs the defined commands with it
 */

#pragma once

#include "Config.h"

//- Declarations -----------------------------------------------------------------------------------

void action_bun( bool, char *[], char *[] );
void action_run( bool, char *[] );
void action_test( bool );
void action_clean( bool );
void action_debug( bool );
void action_config( bool );
void action_compile( bool, char *[] );

