import std.stdio;
void main() {
   // example 1
   auto b1 = ! false;
   // example 2
   auto b2 = false || true;
   // example 3
   auto b3 = true && true;
   // print
   [b1, b2, b3].writeln;
}
