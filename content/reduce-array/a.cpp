#include <iostream>
#include <numeric>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   // example A
   auto sA = std::accumulate(a.begin(), a.end(), std::string());
   // example B
   auto f = [](std::string sY, std::string sZ) {
      return sY + sZ;
   };
   auto sB = std::accumulate(a.begin(), a.end(), std::string(), f);
   // print
   std::cout << (sA == "MayJune" && sB == "MayJune") << std::endl;
}
