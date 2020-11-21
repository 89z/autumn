import std.stdio;

void main() {
   auto n = 0x21 + 1_000;
   auto n2 = 1e3;
   writeln(n == 1033 && n2 == 1000);
}
