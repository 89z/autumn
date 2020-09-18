import std.path, std.stdio;

void main() {
   auto sA = "index.md";
   auto sB = sA.absolutePath;
   sB.writeln;
}
