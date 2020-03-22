#!/bin/dash
set -e -u
x86_64-w64-mingw32-gcc -v test.c \
-DCURL_STATICLIB \
-I /Path/include \
-L /Path/lib \
-lcurl \
-lwldap32 \
-lws2_32
./a
rm a.exe
