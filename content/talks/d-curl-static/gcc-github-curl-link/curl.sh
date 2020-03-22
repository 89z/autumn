#!/bin/dash
set -e -u

if test -d curl
then
   cd curl
   make -C lib -f Makefile.m32 CROSSPREFIX=x86_64-w64-mingw32- clean
   make \
   -C lib \
   -f Makefile.m32 \
   -j 5 \
   CROSSPREFIX=x86_64-w64-mingw32-
else
   git clone --depth 1 git://github.com/curl/curl
   cd curl
   make \
   -C lib \
   -f Makefile.m32 \
   -j 5 \
   CROSSPREFIX=x86_64-w64-mingw32-
fi

mkdir -p /Path/include/curl /Path/lib
cp include/curl/*.h /Path/include/curl
cp lib/libcurl.a /Path/lib
