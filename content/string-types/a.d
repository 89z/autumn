import std.stdio;

void main() {
   // example 1
   string s1 = "One";
   // example 2
   auto s2 = "Two";
   // example 3
   auto s3 = "Three
Three";
   // example 4
   auto s4 = `Four\Four
Four`;
   // example 5
   auto s5 = q{FiveFive
Five};
   // print
   writeln(s1, s2, s3, s4, s5);
}
