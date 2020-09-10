import std.stdio;

ulong f(string s) {
   return s.length;
}

void main() {
   auto n = f("May");
   writeln(n == 3);
}
