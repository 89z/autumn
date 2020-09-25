#include <iostream>

int main() {
   std::string s = "June";
   // example 1
   auto s1 = s[0];
   // example 2
   auto s2 = s.front();
   // example 3
   auto s3 = s.at(0);
   // example 4
   auto s4 = s.substr(0, 2);
   // print
   std::cout << (
      s1 == 'J' && s2 == 'J' && s3 == 'J' && s4 == "Ju"
   ) << std::endl;
}
