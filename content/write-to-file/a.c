#include <stdio.h>
int main() {
   FILE *f = fopen("a.txt", "w");
   fputs("Sunday\n", f);
}
