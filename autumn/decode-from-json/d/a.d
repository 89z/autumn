import std.json, std.stdio;

void main() {
   auto s = `{"month": 12, "day": 31}`;
   auto m = s.parseJSON;
   auto n = m["day"].integer;
   writeln(n == 31);
}
