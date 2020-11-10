import std.stdio;

void main() {
   // example 1
   auto s1 = r"March
April";
   // example 2
   auto s2 = `March
April`;
   // print
   [s1, s2].writeln;
}
