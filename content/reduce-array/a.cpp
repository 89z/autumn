#include <iostream>
#include <numeric>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   // example 1
   auto s1 = std::accumulate(a.begin(), a.end(), std::string());
   // example 2
   auto f = [](std::string s, std::string s2) {
      return s + s2;
   };
   auto s2 = std::accumulate(a.begin(), a.end(), std::string(), f);
   // print
   std::cout << (s1 == "MayJune" && s1 == "MayJune") << std::endl;
}
