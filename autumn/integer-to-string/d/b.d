import std.conv, std.stdio;

void main() {
   auto n = 11;
   auto s = to!string(n);
   writeln(s == "11");
}
