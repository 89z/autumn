#include <filesystem>
namespace fs = std::filesystem;

int main() {
   fs::remove("a.txt");
}
