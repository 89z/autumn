import
   std.conv,
   std.stdio;

void main() {
   auto s = "9.9";
   auto n = s.to!float;
   writeln(n == 9.9f);
}
