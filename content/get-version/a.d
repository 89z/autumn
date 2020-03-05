import std.compiler, std.stdio;
void main() {
   auto n1 = version_major;
   auto n2 = version_minor;
   [n1, n2].writeln;
}
