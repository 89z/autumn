import std.file, std.stdio;
void main() {
   auto b = "index.md".exists;
   b.writeln;
}
