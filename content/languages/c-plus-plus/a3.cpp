#include <iostream>
int main() {
   auto b1 = true;
   // example 1
   std::cout << b1 << std::endl;
   // example 2
   std::cout << std::boolalpha << b1 << std::endl;
   // example 3
   std::cout << std::noboolalpha << b1 << std::endl;
}
