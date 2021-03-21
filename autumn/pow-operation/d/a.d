import std.math, std.stdio;

void main() {
   { // example 1
      float n = pow(10, 5);
      writeln(n == 1e5);
   }
   { // example 2
      float n = pow(9, 0.5);
      writeln(n == 3);
   }
}
