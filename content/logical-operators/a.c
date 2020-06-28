#include <stdio.h>
int main() {
   // example 1
   int b1 = ! 0;
   // example 2
   int b2 = 0 || 1;
   // example 3
   int b3 = 1 && 1;
   // print
   printf("%d %d %d\n", b1, b2, b3);
}
