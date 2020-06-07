git clone --depth 1 git://github.com/curl/curl
Set-Location curl/lib
mingw32-make `
-f Makefile.m32 `
-j 5 `
CFG=-winssl `
CURL_CFLAG_EXTRAS=-fno-addrsig
Copy-Item libcurl.a ..\..\curl.lib
