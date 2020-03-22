#!/bin/dash
cp /Path/lib/libcurl.a curl.lib
# lld-link: error: curl.lib(easy.o): invalid symbol index in addrsig section
ldc2 test.d curl.lib wldap32.lib
rm curl.lib test.lib test.obj
