import std.stdio;

void main() {
   auto m = ["month": 5, "day": 4];
   auto n = m["day"];
   writeln(n == 4);
}
