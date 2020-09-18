#include <iostream>

int main() {
   // example 1
   auto f1 = [](int n, int n1) -> bool {
      return n > n1;
   };
   // example 2
   auto f2 = [](int n, int n2) {
      return n > n2;
   };
   // print
   auto b1 = f1(9, 8);
   auto b2 = f2(9, 8);
   std::cout << (b1 && b2) << std::endl;
}
