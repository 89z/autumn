import std.algorithm.searching, std.stdio;
void main() {
   auto s = "Sunday";
   auto b = endsWith(s, "ay");
   b.writeln;
}
