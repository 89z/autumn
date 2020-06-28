#include <stdio.h>
#include <stdlib.h>
int main() {
   char *s = getenv("BROWSER");
   puts(s);
}
