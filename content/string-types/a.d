import std.stdio;

void main() {
   // example 1
   auto s1 = "May
June";
   // example 2
   auto s2 = `May
June`;
   // example 3
   auto s3 = q{May
June};
   // print
   [s1, s2, s3].writeln;
}
