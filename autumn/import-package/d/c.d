import std.json, std.stdio;

void main() {
   auto m = ["month": 12, "day": 31].JSONValue;
   auto s = m.toJSON;
   s.writeln;
}
