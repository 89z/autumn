import std.file, std.stdio;
void main() {
   auto a1 = dirEntries(".", SpanMode.shallow);
   a1.writeln;
}
