import std.format, std.stdio;
void main() {
   auto n1 = 1000;
   auto s1 = "%5d".format(n1);
   auto s2 = "%,d".format(n1);
   [s1, s2].writeln;
}
