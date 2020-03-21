#!/bin/dash
echo clang
clang -c hello.c
echo llvm-ar
llvm-ar r hello.lib hello.o
echo dmd
/Root/dmd2/windows/bin64/dmd hello_test.d hello.lib
rm hello.lib hello.o hello_test.obj
