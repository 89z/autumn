#include <iostream>

int main() {
   // example 1
   auto n = 10;
   auto s1 = std::to_string(n);
   // example 2
   std::string s2 = std::to_string(7654321);
   int n2 = s2.length() - 3;
   while (n2 > 0) {
      s2.insert(n2, ",");
      n2 -= 3;
   }
   // print
   std::cout << (s1 == "10" && s2 == "7,654,321") << std::endl;
}
