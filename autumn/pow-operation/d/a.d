import io = std.stdio;
import math = std.math;

void main() {
   // example 1
   float n1 = math.pow(10, 5);
   // example 2
   float n2 = math.pow(9, 0.5);
   // print
   io.writeln(n1 == 1e5 && n2 == 3);
}
