#include <stdio.h>
#include <windows.h>
int main(int argc, char* argv[]) {
   char *s1;
   for (int n1 = 0; n1 < argc; n1++) {
      s1 = argv[n1];
      printf("%d %s\n", n1, s1);
   }
   int n2;
   LPWSTR s2;
   LPWSTR *a1 = CommandLineToArgvW(GetCommandLineW(), &n2);
   for (int n1 = 0; n1 < n2; n1++) {
      s2 = a1[n1];
      printf("%d %ws\n", n1, s2);
   }
}
