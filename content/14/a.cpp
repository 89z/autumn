#include <iostream>
#include <vector>

int main() {
   std::vector<int> a = {10, 11};
   for (auto n: a) {
      std::cout << n << std::endl;
   }
}
