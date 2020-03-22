#!/bin/dash
clang -c hello.c
llvm-ar r hello.a hello.o
clang test.c hello.a
./a
rm a.exe hello.a hello.o
