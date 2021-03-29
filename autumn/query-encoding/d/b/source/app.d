import std.stdio, urllibparse;

void main() {
   auto m = ["west": "left", "east": "right"];
   auto s = m.urlEncode;
   s.writeln;
}
