#include <iostream>
#include <map>
#include <string>
int main() {
   std::map<std::string, int> m1 = {{"Sun", 10}, {"Mon", 11}};
   for (auto &[s1, n1]: m1) {
      std::cout << s1 << "," << n1 << std::endl;
   }
}
