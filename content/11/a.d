import std.stdio;
void main() {
   auto o = "a.txt".File;
   foreach (s; o.byLine) {
      s.writeln;
   }
}
