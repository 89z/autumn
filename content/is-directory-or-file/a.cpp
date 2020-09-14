#include <iostream>
#include <unistd.h>

int main() {
   int n = access("index.md", F_OK);
   std::cout << (n != -1) << std::endl;
}
