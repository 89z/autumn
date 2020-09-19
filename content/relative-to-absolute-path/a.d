import std.path, std.stdio;

void main() {
   auto s = "index.md";
   auto s1 = s.absolutePath;
   s1.writeln;
}
