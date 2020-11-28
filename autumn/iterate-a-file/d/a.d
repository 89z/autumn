import std.stdio;

void main() {
   auto o = File("a.d");
   foreach (s; o.byLine) {
      s.writeln;
   }
}
