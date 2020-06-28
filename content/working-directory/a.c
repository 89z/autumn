#include <stdio.h> // puts
#include <stdlib.h> // PATH_MAX
#include <unistd.h> // getcwd
int main() {
   char s[PATH_MAX];
   getcwd(s, PATH_MAX);
   puts(s);
}
