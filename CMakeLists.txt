
cmake_minimum_required( VERSION 3.30 )

project( dev VERSION 1.0 LANGUAGES C )

include_directories( include )
add_subdirectory( src )

set( CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/build )

