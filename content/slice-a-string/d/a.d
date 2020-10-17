import std.stdio;

void main() {
   auto s = "March";
   // example 1
   auto s1 = s[2];
   // example 2
   auto s2 = s[2 .. 3];
   // print
   writeln(s1 == 'r' && s2 == "r");
}
