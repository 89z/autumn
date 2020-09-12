import std.stdio;

void main() {
   auto m = ["year": 2019];
   auto o = "year" in m;
   writeln(o != null);
}
