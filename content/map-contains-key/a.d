import std.stdio;
void main() {
   auto m = ["year": 2020];
   auto o = "year" in m;
   writeln(o != null);
}
