#include <filesystem>
#include <iostream>
using namespace std::filesystem;

int main() {
   auto s = "C:\\Windows\\write.exe";
   auto s1 = path(s).filename();
   std::cout << (s1 == "write.exe") << std::endl;
}
