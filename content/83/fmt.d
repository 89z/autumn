import std.format, std.stdio;

void main() {
   auto n = 1000;
   auto s = format("%,d", n);
   writeln(s == "1,000");
}
