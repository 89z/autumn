#include <iostream>

int main() {
   std::string s = "June";
   // example 1
   auto s1 = s[0];
   // example 2
   auto s2 = s.at(0);
   // example 3
   auto s3 = s.substr(0, 2);
   // print
   std::cout << (s1 == 'J' && s2 == 'J' && s3 == "Ju") << std::endl;
}
