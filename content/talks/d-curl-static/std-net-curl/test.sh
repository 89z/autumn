#!/bin/dash
rm -f /Root/dmd2/windows/bin64/libcurl.dll
cp /Path/curl/libcurl.a curl.lib
dmd test.d -m64 curl.lib
./test # Failed to load curl, tried "libcurl.dll", "curl.dll".
rm curl.lib test.exe test.obj
