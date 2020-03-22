#!/bin/dash
cp /Path/lib/libcurl.a curl.lib
ldc2 -L/OPT:NOICF test.d curl.lib wldap32.lib
./test.exe
rm curl.lib test.exe test.lib test.obj
