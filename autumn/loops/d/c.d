import std.stdio;

void main() {
   int n;
   writeln("example 1");
   n = 10;
   while (n < 20) {
      writeln(n);
      n++;
   }
   writeln("example 2");
   n = 10;
   while (true) {
      if (n > 19) {
         break;
      }
      writeln(n);
      n++;
   }
}
