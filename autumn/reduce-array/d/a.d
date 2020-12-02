import
   algo = std.algorithm,
   io = std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto n = algo.reduce!((n, s) => n + s.length)(size_t(0), a);
   // example 2
   auto s = algo.reduce!((s, s2) => s ~ s2)(a);
   // print
   io.writeln(n == 7 && s == "MayJune");
}
