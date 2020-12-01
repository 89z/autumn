import std.stdio;

void main() {
   writeln("example 1");
   auto n1 = 10;
   while (n1 < 20) {
      writeln(n1);
      n1++;
   }

   writeln("example 2");
   auto n2 = 10;
   while (true) {
      if (n2 > 19) {
         break;
      }
      writeln(n2);
      n2++;
   }
}
