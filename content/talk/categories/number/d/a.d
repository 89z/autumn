import std.math, std.stdio;

void main() {
   ulong nX = 7;
   ulong nY = 19;
   // example 1
   auto n1 = nX ^^ nY;
   // example 2
   auto n2 = pow(nX, nY);
   // print
   writeln(n1, ',', n2);
}
