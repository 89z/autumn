#include <filesystem>
#include <iostream>
int main() {
   auto s1 = "C:\\Windows\\write.exe";
   auto o1 = std::filesystem::path(s1);
   auto s2 = o1.filename();
   std::cout << s2 << std::endl;
}
