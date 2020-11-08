import std.stdio;

void main() {
   // example 1
   auto f1 = function bool (int n, int n1) {
      return n > n1;
   };
   // example 2
   auto f2 = (int n, int n1) {
      return n > n1;
   };
   // example 3
   auto f3 = function bool (int n, int n1) => n > n1;
   // example 4
   auto f4 = (int n, int n1) => n > n1;
   // print
   auto b1 = f1(9, 8);
   auto b2 = f2(9, 8);
   auto b3 = f3(9, 8);
   auto b4 = f4(9, 8);
   writeln(b1 && b2 && b3 && b4);
}
