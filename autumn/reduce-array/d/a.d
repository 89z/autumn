import
   std.algorithm,
   std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto n = reduce!((n, s) => n + s.length)(size_t(0), a);
   // example 2
   auto s = a.reduce!((s, s2) => s ~ s2);
   // print
   writeln(n == 7 && s == "MayJune");
}
