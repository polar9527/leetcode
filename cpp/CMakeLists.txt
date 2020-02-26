cmake_minimum_required(VERSION 3.15)

project(leetcode)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
set(CMAKE_BUILD_TYPE Debug)

add_executable(leetcode.12 12.integer-to-roman.cpp)
add_executable(leetcode.15 15.3Sum.cpp)
add_executable(offer_2 offer_2.cpp)

target_include_directories(offer_2 PUBLIC
        "${PROJECT_BINARY_DIR}"
        )