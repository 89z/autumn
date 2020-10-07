import std.stdio;

void main() {
   // example 1
   auto s1 = r"May
June";
   // example 2
   auto s2 = `May
June`;
   // print
   [s1, s2].writeln;
}
