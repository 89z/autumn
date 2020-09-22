#include <filesystem>
#include <iostream>
using namespace std::filesystem;

int main() {
   auto s = R"(C:\Windows\notepad.exe)";
   auto s1 = path(s).filename();
   std::cout << (s1 == "notepad.exe") << std::endl;
}
