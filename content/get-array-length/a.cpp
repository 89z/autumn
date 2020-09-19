#include <iostream>
#include <vector>

int main() {
   std::vector<std::string> a = {"May", "June"};
   // example 1
   auto n1 = a.size();
   // example 2
   auto n2 = std::size(a);
   // print
   std::cout << (n1 == 2 && n2 == 2) << std::endl;
}
