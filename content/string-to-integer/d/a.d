import std.conv, std.stdio;

void main() {
   auto s = "11";
   auto n = s.to!int;
   writeln(n == 11);
}
