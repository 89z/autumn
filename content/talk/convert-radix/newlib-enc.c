#include <stdio.h>
// head
char *r64_encode(long n1) {
   char *s_safe = "0123456789abcdefghijklmnopqrstuvwxyz"
   "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
   static char s1[7];
   char *s2 = s1;
   int n2;
   for (n2 = 0; n2 < 6; n2++) {
      if (n1 == 0) {
         *s2 = 0;
      }
      *s2++ = s_safe[n1 & 63];
      n1 >>= 6;
   }
   return s1;
}
// body
int main() {
   char *s1 = r64_encode(1234567890);
   puts(s1);
}
