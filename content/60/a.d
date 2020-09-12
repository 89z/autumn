import std.conv, std.stdio;

void main() {
   auto n = 10;
   auto s = n.to!string(16);
   writeln(s);
}
