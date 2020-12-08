Set-Location C:\msys2\mingw64\ssl
Invoke-WebRequest -OutFile cert.pem https://curl.haxx.se/ca/cacert.pem
