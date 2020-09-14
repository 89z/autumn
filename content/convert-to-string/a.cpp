#include <iostream>
#include <string>

int main() {
   auto n = 9;
   auto s = std::to_string(n);
   std::cout << (s == "9") << std::endl;
}
