#include <cstdlib>
#include <iostream>
#include <fstream>
#include <filesystem>
namespace fs = std::filesystem;
 
int main()
{
    fs::create_directories("sandbox/dir/subdir");
    std::ofstream("sandbox/file1.txt").put('a');
    fs::copy("sandbox/file1.txt", "sandbox/file2.txt"); // copy file
    fs::copy("sandbox/dir", "sandbox/dir2"); // copy directory (non-recursive)
    const auto copyOptions = fs::copy_options::update_existing
                           | fs::copy_options::recursive
                           | fs::copy_options::directories_only
                           ;
    fs::copy("sandbox", "sandbox_copy", copyOptions); 
    static_cast<void>(std::system("tree"));
    fs::remove_all("sandbox");
    fs::remove_all("sandbox_copy");
}
