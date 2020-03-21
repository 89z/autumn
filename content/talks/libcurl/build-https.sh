#!/bin/dash
set -e -u
git clone --depth 1 git://github.com/curl/curl
cd curl
mingw32-make -C lib -f Makefile.m32 -j 5 CFG=-winssl
mkdir -p /Path/llvm-mingw/lib /Path/llvm-mingw/include/curl
cp lib/libcurl.a /Path/llvm-mingw/lib
cp include/curl/*.h /Path/llvm-mingw/include/curl
