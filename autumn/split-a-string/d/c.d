import std.stdio;

void main() {
   auto s = "north";
   auto b = cast(ubyte[]) s;
   b.writeln;
}
