import std.json, std.stdio;
void main() {
   auto s = `{"Sunday": 10}`;
   auto m = s.parseJSON;
   m.writeln;
}
