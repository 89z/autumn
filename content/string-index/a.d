import std.algorithm, std.stdio;
void main() {
   auto s = "Sunday";
   auto n = s.countUntil("day");
   writeln(n == 3);
}
