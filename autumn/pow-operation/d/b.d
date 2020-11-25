import std.stdio;

void main() {
   // example 1
   auto n1 = 10 ^^ 5;
   // example 2
   auto n2 = 9 ^^ .5;
   // print
   writeln(n1 == 1e5 && n2 == 3);
}
