import
   algo = std.algorithm,
   io = std.stdio;

void main() {
   // example 1
   auto a1 = ["May", "June"];
   algo.sort(a1);
   // example 2
   auto a2 = [10, 9];
   algo.sort(a2);
   // print
   io.writeln(a1, a2);
}
