#include <stdio.h>
#include <string.h>
int main() {
   char s1[] = "Sunday";
   char s2[3];
   strncpy_s(s2, 3, s1 + strlen(s1) - 2, 2);
   puts(s2);
}
