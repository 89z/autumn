import std.stdio;

void main() {
   auto s = "March";
   auto a = cast(ubyte[]) s;
   a.writeln;
}
