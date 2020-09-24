#include <iostream>

int main() {
   auto n = 10;
   // example 1
   std::cout << n << std::endl;
   // example 2
   std::printf("%d\n", n);
}
   std::string s2 = std::to_string(7654321);
   int n2 = s2.length() - 3;
   while (n2 > 0) {
      s2.insert(n2, ",");
      n2 -= 3;
   }
