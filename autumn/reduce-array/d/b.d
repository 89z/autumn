import
   algo = std.algorithm,
   io = std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto n = algo.fold!((n, s) => n + s.length)(a, size_t(0));
   // example 2
   auto s = algo.fold!((s, s2) => s ~ s2)(a);
   // print
   io.writeln(n == 7 && s == "MayJune");
}
