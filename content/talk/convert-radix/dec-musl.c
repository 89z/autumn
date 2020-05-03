#include <stdio.h>
#include <string.h>
// begin
int r64_decode(char *s_in) {
   char s_safe[] =
   "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_";
   int n_out = 0;
   int n_sh = 0;
   while (*s_in != 0) {
      char *s_out = strchr(s_safe, *s_in);
      n_out |= (s_out - s_safe) << n_sh;
      s_in += 1;
      n_sh += 6;
   }
   return n_out;
}
// end
int main() {
   int n1 = r64_decode("ibwB91");
   printf("%d\n", n1);
}
