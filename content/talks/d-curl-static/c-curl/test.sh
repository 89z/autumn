#!/bin/dash
set -e -u
clang \
-DCURL_STATICLIB \
-I C:/Path/include \
-L C:/Path/lib \
-lcurl \
-lwldap32 \
-lws2_32 \
test.c
./a
rm a.exe
