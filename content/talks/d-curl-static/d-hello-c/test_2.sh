#!/bin/dash
clang -c hello.c
llvm-ar r hello.lib hello.o
dmd test.d hello.lib # Error 43: Not a Valid Library File
rm hello.lib hello.o test.obj
