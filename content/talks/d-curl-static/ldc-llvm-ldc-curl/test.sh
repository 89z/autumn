#!/bin/dash
set -e -u
ldc2 test.d C:/Path/lib/curl_a.lib crypt32.lib wldap32.lib
./test.exe
rm test.exe test.obj
