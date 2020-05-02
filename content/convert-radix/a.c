#include <stdint.h>
#include <stdio.h>
char *radix64(int n1) {
   char *digits = "0123456789abcdefghijklmnopqrstuvwxyz"
   "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
   static char s1[7];
   char *s2;
   s2 = s1;
   while (n1 > 0) {
      *s2++ = digits[n1 & 63];
      n1 >>= 6;
   }
   *s2 = 0;
   return s1;
}
int main() {
   int n1 = 1588337932;
   char *s1 = radix64(n1);
   puts(s1);
}
