import std.stdio;

void main() {
   auto s = "May";
   auto o = File("a.txt", "w");
   o.writeln(s);
}
