#include <iostream>

int main() {
   int n = 7654321;

   // example 1
   std::string s1 = std::to_string(n);
   int n1 = s1.length() - 3;
   while (n1 > 0) {
      s1.insert(n1, ",");
      n1 -= 3;
   }

   // example 2
   std::string s = std::to_string(n);
   std::string s2;
   int n2 = s.length();
   for (auto s2a: s) {
      if (n2 % 3 == 0) {
         s2 += ",";
      }
      s2 += s2a;
      n2--;
   }

   // print
   std::cout << (s1 == "7,654,321" && s2 == "7,654,321") << std::endl;
}
