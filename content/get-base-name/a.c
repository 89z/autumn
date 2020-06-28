#include <libgen.h>
#include <stdio.h>
int main() {
   char s1[] = "/sun/mon.tar.xz";
   char *s2 = basename(s1);
   puts(s2);
}
