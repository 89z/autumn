#include <iostream>
#include <numeric>
#include <vector>

std::string f(std::string s_acc, std::string s_cur) {
   return s_acc + s_cur;
}

int main() {
   std::vector<std::string> a = {"May", "June"};
   auto s = std::accumulate(a.begin(), a.end(), std::string(), f);
   std::cout << s << std::endl;
}
