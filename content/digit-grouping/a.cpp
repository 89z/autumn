#include <iostream>

int main() {
   int n = 7654321;

   // example 1
   std::string s1 = std::to_string(n);
   int n1a = s1.length() - 3;
   while (n1a > 0) {
      s1.insert(n1a, ",");
      n1a -= 3;
   }

   // example 2
   std::string s2a = std::to_string(n);
   std::string s2;
   int n2a = s2a.length();
   for (auto s2b: s2a) {
      if (n2a % 3 == 0) {
         s2 += ",";
      }
      s2 += s2b;
      n2a--;
   }

   // print
   std::cout << (s1 == "7,654,321" && s2 == "7,654,321") << std::endl;
}
