import std.stdio;

void main() {
   auto a = ["west", "east"];
   // example 1
   foreach (s; a) {
      s.writeln;
   }
   // example 2
   foreach (n, s; a) {
      writeln(n, ',', s);
   }
}
