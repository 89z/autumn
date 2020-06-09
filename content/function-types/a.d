import std.stdio;
ulong f(string s) {
   return s.length;
}
void main() {
   auto n = "Sunday".f;
   n.writeln;
}
