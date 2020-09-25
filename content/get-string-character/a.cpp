#include <iostream>

int main() {
   std::string s = "May";
   // example 1
   auto s1 = s.front();
   // example 2
   auto s2 = s.at(1);
   // example 3
   auto s3 = s[1];
   // example 4
   auto s4 = s.substr(1, 1);
   // print
   std::printf("%d\n", s1 == 'M' && s2 == 'a' && s3 == 'a' && s4 == "a");
}
