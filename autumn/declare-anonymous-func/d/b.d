import std.stdio;

void main() {
   { // example 1
      auto f = function int (int d, int e) => d + e;
      writeln(f(4, 5) == 9);
   }
   { // example 2
      auto f = (int d, int e) => d + e;
      writeln(f(4, 5) == 9);
   }
}
