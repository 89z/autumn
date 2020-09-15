#include <iostream>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   std::string s;
   for (auto s2: a) {
      if (s != "") {
         s += ",";
      }
      s += s2;
   }
   std::cout << s << std::endl;
}
