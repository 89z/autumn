#include <iostream>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   std::string s;
   for (auto s1: a) {
      if (s != "") {
         s += ",";
      }
      s += s1;
   }
   std::cout << s << std::endl;
}
