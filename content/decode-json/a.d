import std.json, std.stdio;
void main() {
   auto s = `{"year": 2019, "month": 12, "day": 31}`;
   auto m = s.parseJSON;
   m.writeln;
}
