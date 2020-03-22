#!/bin/dash
cp /Path/lib/libcurl.a curl.lib
ldc2 -c test.d
lld-link \
/DEFAULTLIB:vcruntime140 \
/LIBPATH:C:/Root/ldc2/lib \
/LIBPATH:C:/Root/ldc2/lib/mingw \
/OPT:NOICF \
curl.lib \
druntime-ldc.lib \
ldc_rt.builtins.lib \
phobos2-ldc.lib \
test.obj \
wldap32.lib
./curl.exe
rm -f curl.exe curl.lib test.lib test.obj
