#include <iomanip>
#include <iostream>

int main() {
   std::time_t n = 1577858399;
   auto o = std::localtime(&n);
   auto s = std::put_time(o, "%c");
   std::cout << s << std::endl;
}
