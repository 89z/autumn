import std.stdio;

void main() {
   // example 1
   auto s1 = "One";
   // example 2
   auto s2 = "Two
Two";
   // example 3
   auto s3 = `Three
Three\`;
   // example 4
   auto s4 = q{Four
Four};
   // print
   writeln(s1, s2, s3, s4);
}
