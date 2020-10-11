import std.file, std.stdio;

void main() {
   "..".chdir;
   auto s = getcwd;
   s.writeln;
}
