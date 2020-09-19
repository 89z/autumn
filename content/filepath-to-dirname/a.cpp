#include <filesystem>
#include <iostream>
namespace fs = std::filesystem;

int main() {
   fs::path o = "C:\\Windows\\notepad.exe";
   auto s = o.parent_path();
   std::cout << (s == "C:\\Windows") << std::endl;
}
