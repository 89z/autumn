#include <iostream>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   for (auto s: a) {
      std::cout << s << std::endl;
   }
}
