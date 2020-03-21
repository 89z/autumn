#!/bin/dash
set -e -u
clang++ \
-DCURL_STATICLIB \
-I C:/Path/llvm-mingw/include \
-L C:/Path/llvm-mingw/lib \
-lcurl \
-lwldap32 \
-lws2_32 \
http.cpp
