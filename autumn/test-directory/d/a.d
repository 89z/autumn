import
   std.file,
   std.stdio;

bool testDir(string s) {
   return s.exists && s.isDir;
}

void main() {
   auto b = testDir(`C:\Users`);
   b.writeln;
}
