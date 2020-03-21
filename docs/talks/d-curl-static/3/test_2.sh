#!/bin/dash
dmd -lib hello.d
dmd test.d hello.lib
./test
rm hello.lib test.exe test.obj
