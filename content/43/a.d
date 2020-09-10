import std.algorithm.searching, std.stdio;

void main() {
   auto s = "June";
   auto b = endsWith(s, "ne");
   b.writeln;
}
