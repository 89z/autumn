import std.algorithm, std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto s = a.reduce!((sa, sc) => sa ~ sc);
   // example 2
   auto f = (string sa, string sc) => sa ~ sc;
   auto s2 = a.reduce!(f);
   // print
   writeln(s == "MayJune" && s2 == "MayJune");
}
