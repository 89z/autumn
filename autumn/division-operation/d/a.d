import std.stdio;

void main() {
   auto f = 7.5, i = 7;
   // example 1
   auto n1 = f / 2;
   // example 2
   auto n2 = cast(int) f / 2;
   // example 3
   auto n3 = i / 2;
   // example 4
   auto n4 = cast(real) i / 2;
   // print
   writeln(n1 == 3.75 && n2 == 3 && n3 == 3 && n4 == 3.5);
}
