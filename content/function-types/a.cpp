#include <iostream>

bool f(int n, int n2) {
   return n > n2;
}

int main() {
   auto b = f(9, 8);
   std::cout << b << std::endl;
}
