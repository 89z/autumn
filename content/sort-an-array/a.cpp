#include <algorithm>
#include <iostream>

int main() {
   // example 1
   std::vector<std::string> a1 = {"May", "June"};
   std::sort(a1.begin(), a1.end());
   for (auto s: a1) {
      std::cout << s << std::endl;
   }
   // example 2
   std::vector<int> a2 = {10, 9};
   std::sort(a2.begin(), a2.end());
   for (auto n: a2) {
      std::cout << n << std::endl;
   }
}
