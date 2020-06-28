#include <stdio.h>
#include <string.h>
int main() {
   char *s1 = "Sunday";
   char *s2 = strstr(s1, "un");
   printf("%d\n", s2 != NULL);
}
