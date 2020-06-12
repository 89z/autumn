import std.json, std.stdio;
void main() {
   auto m = ["Sunday": 10];
   auto o = m.JSONValue;
   auto s = o.toJSON;
   s.writeln;
}
