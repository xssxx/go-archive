#!/bin/bash

# output
TARGET_C="target/ll_c"
TARGET_CPP="target/ll_cpp"

mkdir -p target

# Compile the C program
gcc linkedList.c -o $TARGET_C
if [ $? -eq 0 ]; then
    echo "C Compilation successful. Running the C program..."
    # Run the C program
    ./$TARGET_C
else
    echo "C Compilation failed."
fi

# Compile the C++ program
g++ linkedList.cpp -o $TARGET_CPP
if [ $? -eq 0 ]; then
    echo "C++ Compilation successful. Running the C++ program..."
    # Run the C++ program
    ./$TARGET_CPP
else
    echo "C++ Compilation failed."
fi
