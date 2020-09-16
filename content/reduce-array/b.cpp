#include <iostream>
#include <vector>

std::string f(std::string s_acc, std::string s_cur) {
   return s_acc + s_cur;
}

int main() {
   std::vector<std::string> a = {"May", "June"};
   std::string s;
   for (auto s_cur: a) {
      s = f(s, s_cur);
   }
   std::cout << (s == "MayJune") << std::endl;
}
