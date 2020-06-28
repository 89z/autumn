#include <stdio.h>
int main() {
   int n = 10;
   char *s;
   if (n > 12) {
      s = "Tue";
   } else if (n > 11) {
      s = "Mon";
   } else {
      s = "Sun";
   }
   puts(s);
}
