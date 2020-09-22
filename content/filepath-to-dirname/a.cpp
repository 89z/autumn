#include <filesystem>
#include <iostream>
namespace fs = std::filesystem;

int main() {
   fs::path o = R"(C:\Windows\notepad.exe)";
   auto s = o.parent_path();
   std::cout << (s == R"(C:\Windows)") << std::endl;
}
