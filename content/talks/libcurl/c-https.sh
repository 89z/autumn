#!/bin/dash
set -e -u
clang \
-DCURL_STATICLIB \
-I C:/Path/llvm-mingw/include \
-L C:/Path/llvm-mingw/lib \
-lcrypt32 \
-lcurl \
-lwldap32 \
-lws2_32 \
https.c
