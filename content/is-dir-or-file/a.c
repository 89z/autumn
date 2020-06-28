#include <stdio.h>
#include <sys/stat.h>
int main() {
   int n = access("index.md", F_OK);
   printf("%d\n", n != -1);
}
