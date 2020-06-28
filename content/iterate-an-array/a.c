#include <stdio.h>
int main() {
   char *a[2] = {"Sunday", "Monday"};
   for (int n = 0; n <= 1; n++) {
      char *s = a[n];
      puts(s);
   }
}
