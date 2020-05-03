#include <stdio.h>
// head
char *radix64(int n1) {
   char *s_safe = "0123456789abcdefghijklmnopqrstuvwxyz"
   "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
   static char s1[7];
   int n2;
   for (n2 = 0; n1 > 0; n2++) {
      s1[n2] = s_safe[n1 & 63];
      n1 >>= 6;
   }
   s1[n2] = 0;
   return s1;
}
// body
int main() {
   char *s1 = radix64(1588337932);
   puts(s1);
}
