#include <stdio.h>
#include <unistd.h>
int main() {
   int n1 = fileno(stdin);
   int n2 = isatty(n1);
   printf("%d\n", n2 != 0);
}
