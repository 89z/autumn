import
   std.conv,
   std.stdio;

void main() {
   auto n = 100;
   auto s = to!string(n);
   writeln(s == "100");
}
