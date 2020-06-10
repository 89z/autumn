import std.stdio;
void main() {
   auto m = ["Sunday": 10];
   auto o = "Sunday" in m;
   writeln(o != null);
}
