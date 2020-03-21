#!/bin/dash
clang -c hello.c
llvm-ar r hello.a hello.o
dmd hello_test.d hello.a # Error: unrecognized file extension a
rm hello.a hello.o
