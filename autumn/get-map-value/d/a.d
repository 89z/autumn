import std.stdio;

void main() {
   auto m = ["month": 12, "day": 31];
   auto n = m["day"];
   writeln(n == 31);
}
