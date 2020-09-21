#include <iostream>

int main() {
   // example 1
   auto n1 = 10;
   // example 2
   auto n2 = 1'000;
   // example 3
   auto n3 = 1e6;
   // example 4
   auto n4 = 9.9;
   // print
   std::printf("%d %d %e %f\n", n1, n2, n3, n4);
}
