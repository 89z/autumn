import std.conv, std.stdio;

void main() {
   auto n = 9;
   // example 1
   auto s = n.to!(string);
   // example 2
   auto s2 = n.to!string;
   // example 3
   auto s3 = n.text;
   // print
   [s == "9", s2 == "9", s3 == "9"].writeln;
}
