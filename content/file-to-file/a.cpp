#include <filesystem>
namespace fs = std::filesystem;
int main() {
   fs::copy("a.txt", "b.txt");
}
