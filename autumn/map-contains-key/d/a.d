import std.stdio;

bool key_exists(string s, int[string] m) {
   return (s in m) != null;
}

void main() {
   auto m = ["month": 5, "day": 4];
   auto b = key_exists("day", m);
   b.writeln;
}
