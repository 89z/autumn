#include <iostream>
#include <map>
#include <string>

int main() {
   std::map<std::string, int> m = {{"year", 2020}, {"month", 9}};
   for (auto &[s, n]: m) {
      std::cout << s << ":" << n << std::endl;
   }
}
