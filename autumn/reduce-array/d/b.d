import
   std.algorithm,
   std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto n = a.fold!((n, s) => n + s.length)(size_t(0));
   // example 2
   auto s = a.fold!((s, s2) => s ~ s2);
   // print
   writeln(n == 7 && s == "MayJune");
}
