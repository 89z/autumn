#include <fcntl.h>
#include <stdio.h>
int wmain(int argc, wchar_t **argv) {
   _setmode(_fileno(stdout), _O_U16TEXT);
   wprintf(L"%s\n", argv[1]);
}
