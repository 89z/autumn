#include <filesystem>
#include <iostream>
namespace fs = std::filesystem;

int main() {
   fs::path o = "C:\\Windows\\write.exe";
   auto s = o.parent_path();
   std::cout << s << std::endl;
}
