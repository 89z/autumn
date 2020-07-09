#include <iostream>
#include <vector>
int main() {
   std::vector<std::string> a = {"BBBB", "AA", "CCC"};
   std::sort(a.begin(), a.end());
   for (auto s: a) {
      std::cout << s << std::endl;
   }
}
