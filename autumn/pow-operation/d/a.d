import std.math: pow;
import std.stdio: writeln;

void main() {
   // example 1
   float n1 = pow(10, 5);
   // example 2
   float n2 = pow(9, .5);
   // print
   writeln(n1 == 1e5 && n2 == 3);
}
