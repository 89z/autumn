import std.stdio;

void main() {
   auto m = ["year": 2019];
   auto n = m["year"];
   writeln(n == 2019);
}
