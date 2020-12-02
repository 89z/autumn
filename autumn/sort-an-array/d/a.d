import std.algorithm: sort;
import std.stdio: writeln;

void main() {
   // example 1
   auto a1 = ["May", "June"];
   a1.sort;
   // example 2
   auto a2 = [10, 9];
   a2.sort;
   // print
   writeln(a1, a2);
}
