import std.file: exists, isDir;
import std.stdio: writeln;

bool testDir(string s) {
   return s.exists && s.isDir;
}

void main() {
   auto b = testDir(`C:\Users`);
   b.writeln;
}
