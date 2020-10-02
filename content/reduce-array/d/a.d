import std.algorithm, std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto s1 = a.reduce!((s, s1) => s ~ s1);
   // example 2
   auto f = (string s, string s2) => s ~ s2;
   auto s2 = a.reduce!(f);
   // print
   writeln(s1 == "MayJune" && s2 == "MayJune");
}
