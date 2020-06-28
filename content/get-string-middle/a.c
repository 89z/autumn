#include <stdio.h>
#include <string.h>
int main() {
   char s1[] = "Sunday";
   char s2[3];
   strncpy(s2, s1 + 1, 2);
   puts(s2);
}
