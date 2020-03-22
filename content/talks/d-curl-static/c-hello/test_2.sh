#!/bin/dash
clang -c hello.c
llvm-ar r hello.a hello.o
clang -c test.c
clang test.o hello.a
./a
rm a.exe hello.a hello.o test.o
