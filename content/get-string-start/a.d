import std.stdio;

void main() {
   auto s = "June";
   // example 1
   auto s2 = s[0];
   // example 2
   auto s3 = s[0 .. 2];
   // print
   [s2 == 'J', s3 == "Ju"].writeln;
}
