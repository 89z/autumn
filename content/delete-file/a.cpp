#include <iostream>
#include <cstdint>
#include <filesystem>
namespace fs = std::filesystem;
int main()
{
    fs::path dir = fs::temp_directory_path();
    fs::create_directories(dir / "abcdef/example");
    std::uintmax_t n = fs::remove_all(dir / "abcdef");
    std::cout << "Deleted " << n << " files or directories\n";
}
