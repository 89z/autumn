import std.stdio;

void main() {
   // example 1
   auto s1 = "May June";
   // example 2
   auto s2 = q{May June};
   // example 3
   auto s3 = `May\June`;
   // print
   writeln(s1, "\n", s2, "\n", s3);
}
