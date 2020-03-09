#include <fcntl.h>
#include <io.h>
#include <iostream>
wmain() {
   _setmode(_fileno(stdout), _O_U16TEXT);
   std::wcout << L"☺☺" << std::endl;
   return 0;
}
