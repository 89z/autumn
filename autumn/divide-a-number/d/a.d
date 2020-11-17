import std.stdio;

void main() {
   // example 1
   auto n1 = 46 % 10;
   // example 2
   auto n2 = 46 / 10;
   // example 3
   auto n3 = 46 / 10.;
   // print
   writeln(n1 == 6 && n2 == 4 && n3 == 4.6);
}
