#!/bin/dash
set -e -u

x86_64-w64-mingw32-gcc \
-DCURL_STATICLIB \
-I/Path/include \
-c \
test.c

x86_64-w64-mingw32-ld test.o \
/usr/x86_64-w64-mingw32/sys-root/mingw/lib/../lib/crt2.o \
-L/Path/lib \
-L/usr/lib/gcc/x86_64-w64-mingw32/9.2.0 \
-L/usr/x86_64-w64-mingw32/sys-root/mingw/lib \
-lcurl \
-lwldap32 \
-lws2_32 \
-lmingw32 \
-lmsvcrt \
-ladvapi32 \
-lkernel32 \
-lgcc \
-lmingwex \
-lmsvcrt

./a.exe
rm a.exe test.o
