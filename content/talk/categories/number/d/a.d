import std.math, std.stdio;

void main() {
   ulong nX = 7;
   ulong nY = 19;
   // example 1
   auto n1 = nX ^^ nY;
   // example 2
   auto n2 = pow(nX, nY);
   // example 3
   ulong n3 = 1;
   for (auto n = 0; n < 19; n++) {
      n3 *= 7;
   }
   // print
   writeln(n1, ',', n2, ',', n3);
}
