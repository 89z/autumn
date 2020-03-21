#!/bin/dash
clang -c hello.c
llvm-ar r hello.lib hello.o
dmd test.d -m64 hello.lib
# MSVCR120.dll: cannot open shared object file: No such file or directory
./test
rm hello.lib hello.o test.exe test.obj
