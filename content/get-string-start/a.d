import std.stdio;

void main() {
   auto s = "June";
   // example 1
   auto s1 = s[0];
   // example 2
   auto s2 = s[0 .. 2];
   // print
   writeln(s1 == 'J' && s2 == "Ju");
}
