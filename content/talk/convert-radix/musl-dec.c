#include <stdint.h>
#include <stdlib.h>
#include <string.h>
// head
char digits[] =
"./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
// body
long r64_decode(const char *s) {
   int e;
   uint32_t x = 0;
   for (e = 0; e < 36 && *s; e += 6, s++) {
      const char *d = strchr(digits, *s);
      if (!d) {
         break;
      }
      x |= (uint32_t)(d-digits) << e;
   }
   return (int32_t)(x);
}
