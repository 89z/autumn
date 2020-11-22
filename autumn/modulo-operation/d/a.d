import std.stdio;

void main() {
   auto f = 7.5, i = 7;
   // example 1
   auto n1 = f % 2;
   // example 2
   auto n2 = i % 2;
   // print
   writeln(n1 == 1.5 && n2 == 1);
}
