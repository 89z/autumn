import std.algorithm, std.stdio;

void main() {
   auto s = "June";
   auto n = s.countUntil("un");
   writeln(n == 1);
}
