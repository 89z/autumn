import std.stdio;

void main() {
   auto f = File("a.d");
   foreach (s; f.byLine) {
      s.writeln;
   }
}
