#include <iostream>

int main() {
   // example 1
   auto s1 = "One";
   // example 2
   auto s2 = R"(Two\Two
Two)";
   // print
   std::cout << s1 << s2 << std::endl;
}
