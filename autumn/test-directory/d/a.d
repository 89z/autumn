import std.file, std.stdio;

void main() {
   auto s = `C:\Users`;
   writeln(s.exists && s.isDir);
}
