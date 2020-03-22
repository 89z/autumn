#!/bin/dash
set -e -u
git clone --depth 1 git://github.com/curl/curl
cd curl

make -C lib -f Makefile.m32 -j 5 CROSSPREFIX=x86_64-w64-mingw32- \
CURL_CFLAG_EXTRAS=-D_MSVCRT_

#ifdef _MSVCRT_
#ifdef _UCRT



mkdir -p /Path/lib
cp lib/libcurl.a /Path/lib/curl.lib
