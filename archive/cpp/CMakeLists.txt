cmake_minimum_required(VERSION 3.15)

project(leetcode)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
set(CMAKE_BUILD_TYPE Debug)

add_executable(leetcode.12 12.integer-to-roman.cpp)
add_executable(leetcode.15 15.3Sum.cpp)
add_executable(offer_1 offer_1.cpp)

target_include_directories(offer_1 PUBLIC
        "${PROJECT_BINARY_DIR}"
        )