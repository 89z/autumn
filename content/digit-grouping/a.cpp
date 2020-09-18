#include <iostream>

int main() {
   auto n = 7654321;
   // example 1
   std::string s1 = std::to_string(n);
   int n1 = s1.length() - 3;
   while (n1 > 0) {
      s1.insert(n1, ",");
      n1 -= 3;
   }
   // example 2
   std::string s2a = std::to_string(n);
   std::string s2b;
   int n2a = s2a.length();
   int n2b;
   for (auto s2c: s2a) {
      if (n2a % 3 == n2b % 3) {
         s2b += ",";
      }
      s2b += s2c;
      n2b++;
   }
   // print
   std::cout << (s1 == "7,654,321" && s2b == "7,654,321") << std::endl;
}   
