import std.path, std.stdio;

void main() {
   auto s = extension("a.d");
   writeln(s == ".d");
}
