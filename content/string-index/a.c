#include <stdio.h>
#include <string.h>
int main() {
   char *s1 = "Sunday";
   char *s2 = strstr(s1, "day");
   int n = s2 - s1;
   printf("%d\n", n);
}
