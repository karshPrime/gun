
#pragma once

#include <stdio.h>

typedef enum {
    NONE, RUN, BUN, INIT, TEMPLATE, TEST, DEBUG, LICENSE, CONFIG, HELP, VERSION, COMPILE, CLEAN
} Commands;

void print_help( Commands );
void print_usage( Commands );
void print_version( void );

extern const char *CURRENT_VERSION;

