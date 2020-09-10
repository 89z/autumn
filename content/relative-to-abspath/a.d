import std.path, std.stdio;

void main() {
   auto s = "index.md";
   auto s2 = s.absolutePath;
   s2.writeln;
}
