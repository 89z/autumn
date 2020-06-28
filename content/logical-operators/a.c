#include <stdio.h>
#include <stdbool.h>
int main() {
   // example 1
   int b1 = ! false;
   // example 2
   int b2 = false || true;
   // example 3
   int b3 = true && true;
   // print
   printf("%d %d %d\n", b1, b2, b3);
}
