#include <iostream>

int main() {
   // example 1
   std::string s1 = "One";
   // example 2
   auto s2 = "Two";
   // example 3
   auto s3 = R"(Three\Three
Three)";
   // print
   std::cout << s1 << s2 << s3 << std::endl;
}
