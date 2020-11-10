import std.stdio;

void main() {
   auto s = "March";
   auto o = File("a.txt", "w");
   o.writeln(s);
}
