#include <chrono>
#include <filesystem>
#include <fstream>
#include <iomanip>
#include <iostream>
namespace fs = std::filesystem;
using namespace std::chrono_literals;
int main() {
   fs::path p = fs::current_path() / "example.bin";
   std::ofstream(p.c_str()).put('a'); // create file
   auto ftime = fs::last_write_time(p);
   std::time_t cftime = decltype(ftime)::clock::to_time_t(ftime);
   fs::last_write_time(p, ftime + 1h); // move file write time 1 hour to the future
}
