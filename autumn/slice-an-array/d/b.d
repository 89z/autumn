import std.stdio;

void main() {
   auto a = [0, 10, 20, 30, 40];
   { // example 1
      auto b = a[2 .. 4];
      b.writeln;
   }
   { // example 2
      auto b = a[2 .. $];
      b.writeln;
   }
}
