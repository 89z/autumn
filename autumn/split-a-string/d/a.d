import std.array, std.stdio;

void main() {
   auto s = "west,east";
   auto a = s.split(",");
   a.writeln;
}
