import std.stdio;

void main() {
   foreach (n; 0 .. 10) {
      if (n % 3 == 0) {
         continue;
      }
      n.writeln;
   }
}
