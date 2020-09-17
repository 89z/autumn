#include <iostream>
#include <vector>

std::string f(std::string sa, std::string sc) {
   return sa + sc;
}

int main() {
   std::vector<std::string> a = {"May", "June"};
   std::string s;
   for (auto sc: a) {
      s = f(s, sc);
   }
   std::cout << (s == "MayJune") << std::endl;
}
