import std.conv, std.stdio;

void main() {
   auto s = "100";
   auto n = s.to!int;
   writeln(n == 100);
}
