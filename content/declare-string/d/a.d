import std.stdio;

void main() {
   // example 1
   auto s1 = q"[sigma\tau]";
   // example 2
   auto s2 = q"[["sigma", "tau"]]";
   // print
   writeln(s1, s2);
}
