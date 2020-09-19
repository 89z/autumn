#include <iostream>

int main() {
   std::string s = "📗";
   auto n = s.length();
   std::cout << (n == 4) << std::endl;
}
