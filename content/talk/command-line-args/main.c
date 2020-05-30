#include <stdio.h>
#include <windows.h>
int main() {
   int argc;
   wchar_t **argv_utf16 = CommandLineToArgvW(GetCommandLineW(), &argc);
   char* arg_utf8;
   WideCharToMultiByte(CP_UTF8, 0, argv_utf16[1], -1, arg_utf8, 3, 0, 0);
   printf("%s\n", arg_utf8);
}
