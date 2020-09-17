#include <iostream>
#include <numeric>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   // example 1
   auto s = std::accumulate(a.begin(), a.end(), std::string());
   // example 2
   auto f = [](std::string sa, std::string sc) {
      return sa + sc;
   };
   auto s2 = std::accumulate(a.begin(), a.end(), std::string(), f);
   // print
   std::cout << (s == "MayJune" && s2 == "MayJune") << std::endl;
}
