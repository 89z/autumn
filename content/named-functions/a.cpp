#include <iostream>

bool f(int n, int n1) {
   return n > n1;
}

int main() {
   auto b = f(9, 8);
   std::cout << b << std::endl;
}
