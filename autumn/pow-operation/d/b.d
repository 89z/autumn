import std.stdio;

void main() {
   { // example 1
      auto n = 10 ^^ 5;
      writeln(n == 1e5);
   }
   { // example 2
      auto n = 9 ^^ 0.5;
      writeln(n == 3);
   }
}
