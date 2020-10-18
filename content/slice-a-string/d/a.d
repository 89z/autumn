import std.stdio;

void main() {
   auto s = "June";
   // example 1
   auto s1 = s[2];
   // example 2
   auto s2 = s[2 .. 3];
   // print
   writeln(s1 == 'n' && s2 == "n");
}
