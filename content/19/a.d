import std.stdio;
void main() {
   // example 1
   auto s = "May
June";
   // example 2
   auto s2 = `May
June`;
   // example 3
   auto s3 = q{May
June};
   // print
   [s, s2, s3].writeln;
}
