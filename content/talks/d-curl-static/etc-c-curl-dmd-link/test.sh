#!/bin/dash
cp /Path/lib/libcurl.a curl.lib
dmd -c -m64 test.d
lld-link \
/DEFAULTLIB:advapi32.lib \
/DEFAULTLIB:curl.lib \
/DEFAULTLIB:ucrtbase.lib \
/DEFAULTLIB:wldap32.lib \
/LIBPATH:C:/Root/dmd2/windows/lib64/mingw \
/LIBPATH:C:/Root/dmd2/windows/lib64 \
/OPT:NOICF \
test.obj
./test.exe
rm curl.lib test.exe test.lib test.obj
