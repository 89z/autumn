import std.file: exists, isFile;
import std.stdio: writeln;

bool testFile(string s) {
   return s.exists && s.isFile;
}

void main() {
   auto b = testFile("a.d");
   b.writeln;
}
