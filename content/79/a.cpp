#include <iostream>
#include <map>
#include <string>

int main() {
   std::map<std::string, int> m = {{"year", 2019}, {"month", 12}};
   for (auto &[s, n]: m) {
      std::cout << s << ":" << n << std::endl;
   }
}
