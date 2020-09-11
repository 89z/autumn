#include <filesystem>
#include <iostream>
using namespace std::filesystem;

int main() {
   auto s = "C:\\Windows\\write.exe";
   auto s2 = path(s).filename();
   std::cout << s2 << std::endl;
}
