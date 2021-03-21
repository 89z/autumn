import std.stdio;

void main() {
   auto f = 7.5, i = 7;
   { // example 1
      auto n = f / 2;
      writeln(n == 3.75);
   }
   { // example 2
      auto n = cast(int) f / 2;
      writeln(n == 3);
   }
   { // example 3
      auto n = i / 2;
      writeln(n == 3);
   }
   { // example 4
      auto n = cast(real) i / 2;
      writeln(n == 3.5);
   }
}
