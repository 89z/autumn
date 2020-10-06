import std.stdio;

void main() {
   auto s = "May";
   // example 1
   auto s1 = s[1];
   // example 2
   auto s2 = s[1 .. 2];
   // print
   writeln(s1 == 'a' && s2 == "a");
}
