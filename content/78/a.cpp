#include <iostream>

size_t f(std::string s) {
   return s.length();
}

int main() {
   auto n = f("📗");
   std::cout << n << std::endl;
}
