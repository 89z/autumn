import
   std.file,
   std.stdio;

bool testFile(string s) {
   return s.exists && s.isFile;
}

void main() {
   auto b = testFile("a.d");
   b.writeln;
}
