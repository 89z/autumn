#include <iostream>

int main() {
   // example 1
   std::string s1 = "May";
   // example 2
   auto s2 = "May";
   // example 3
   auto s3 = R"(May
June)";
   // print
   std::cout << s1 << s2 << s3 << std::endl;
}
