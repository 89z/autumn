#include <stdio.h>
#include <string.h>
int f(char *s) {
   return strlen(s);
}
int main() {
   int n = f("Sunday");
   printf("%d\n", n);
}
