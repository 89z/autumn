import std.stdio;
void main() {
   auto m = ["year": 2020];
   auto n = m["year"];
   writeln(n == 2020);
}
