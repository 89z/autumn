---
title: LibCurl
---

So using these:

- <https://curl.haxx.se/windows>
- <https://github.com/mstorsjo/llvm-mingw>

You can link with LibCurl:

~~~
clang \
-DCURL_STATICLIB \
-lbrotlicommon-static \
-lbrotlidec-static \
-lcrypt32 \
-lcrypto \
-lcurl \
-lmsvcrt-os \
-lnghttp2 \
-lssh2 \
-lssl \
-lwldap32 \
-lws2_32 \
-lz \
-static \
simple.c
~~~

but if you build LibCurl:

~~~
git clone --depth 1 git://github.com/curl/curl
cd curl
mingw32-make -C lib -f Makefile.m32 -j 5
mkdir -p /Path/llvm-mingw/lib /Path/llvm-mingw/include/curl
cp lib/libcurl.a /Path/llvm-mingw/lib
cp include/curl/*.h /Path/llvm-mingw/include/curl
~~~

it is simpler to link:

~~~
clang \
-DCURL_STATICLIB \
-lcrypt32 \
-lcurl \
-lwldap32 \
-lws2_32 \
simple.c
~~~
