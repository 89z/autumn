import std.stdio;

void main() {
   { // natural int
      auto n = 7 / 2;
      writeln(n == 3);
   }
   { // natural float
      auto n = 7.0 / 2;
      writeln(n == 3.5);
   }
   { // force int
      auto n = cast(int) 7.0 / 2;
      writeln(n == 3);
   }
   { // force float
      auto n = cast(real) 7 / 2;
      writeln(n == 3.5);
   }
}
