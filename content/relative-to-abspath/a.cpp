#include <filesystem>
#include <iostream>
namespace fs = std::filesystem;
int main() {
   fs::path o = "index.md";
   auto s = fs::absolute(o);
   std::cout << s << std::endl;
}
