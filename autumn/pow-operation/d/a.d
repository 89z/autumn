import
   std.math,
   std.stdio;

void main() {
   // example 1
   float n1 = pow(10, 9);
   // example 2
   float n2 = pow(9, .5);
   // print
   writeln(n1 == 1e9 && n2 == 3);
}
