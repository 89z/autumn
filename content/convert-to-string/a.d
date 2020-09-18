import std.conv, std.stdio;

void main() {
   auto n = 9;
   // example 1
   auto s1 = n.text;
   // example 2
   auto s2 = n.to!string;
   // print
   writeln(s1 == "9" && s2 == "9");
}
