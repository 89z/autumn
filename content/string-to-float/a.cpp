#include <string>
#include <iostream>

int main() {
   auto s = "9.9";
   // example 1
   auto n1 = std::stof(s);
   // example 2
   auto n2 = std::stod(s);
   // print
   std::cout << (n1 == 9.9f && n2 == 9.9) << std::endl;
}
