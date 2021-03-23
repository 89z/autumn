import std.stdio;

void main() {
   bool b;

   // example 1
   b = ! false;
   b.writeln;

   // example 2
   b = false || true;
   b.writeln;

   // example 3
   b = true && true;
   b.writeln;

}
