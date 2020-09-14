import std.conv, std.stdio;

void main() {
   auto s = "9";
   auto n = s.to!int;
   writeln(n == 9);
}
