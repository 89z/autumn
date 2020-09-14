import std.conv, std.stdio;

void main() {
   auto s = "10";
   auto n = s.to!(int);
   writeln(n == 10);
}
