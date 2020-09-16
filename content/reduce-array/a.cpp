#include <iostream>
#include <numeric>
#include <vector>

std::string f(std::string s_acc, std::string s_cur) {
   return s_acc + s_cur;
}

int main() {
   std::vector<std::string> a = {"May", "June"};
   // example 1
   auto s = std::accumulate(a.begin(), a.end(), std::string(), f);
   // example 2
   std::string s2;
   for (auto s_cur: a) {
      s2 += s_cur;
   }
   // print
   std::cout << (s == "MayJune" && s2 == "MayJune") << std::endl;
}
