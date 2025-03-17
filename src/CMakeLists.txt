
add_executable( ${PROJECT_NAME}
    main.c
    Actions.c
    Commands.c
    Config.c
    Parse.c
    Print.c
)

target_include_directories( ${PROJECT_NAME} PUBLIC ${PROJECT_SOURCE_DIR}/include )

