#include <iostream>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   // example 1
   for (auto s: a) {
      std::cout << s << std::endl;
   }
   // example 2
   int n;
   for (auto s: a) {
      std::cout << n << ',' << s << std::endl;
      n++;
   }
}
