#!/bin/dash
cp /Path/lib/libcurl.a curl.lib
dmd -m64 test.d curl.lib advapi32.lib ucrtbase.lib wldap32.lib
./test.exe
rm -f curl.lib test.exe test.lib test.obj
