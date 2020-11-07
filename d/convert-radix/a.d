import std.conv, std.stdio;

void main() {
   // example 1
   auto n = 1577858399;
   auto s1 = n.to!string(36);
   // example 2
   auto s = "Q3EZBZ";
   auto n2 = s.to!int(36);
   // print
   writeln(s1 == s && n2 == n);
}
