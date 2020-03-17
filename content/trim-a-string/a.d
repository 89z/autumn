import std.string, std.stdio;
void main() {
   auto s1 = " Sun Mon ";
   // example 1
   auto s2 = s1.stripLeft;
   // example 2
   auto s3 = s1.stripRight;
   // example 3
   auto s4 = s1.strip;
   // print
   [s2, s3, s4].writeln;
}
