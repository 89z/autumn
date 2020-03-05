import std.stdio;
void main() {
   // example 1
   auto s1 = "sun";
   s1 = s1 ~ " mon";
   // example 2
   auto s2 = "sun";
   s2 ~= " mon";
   // print
   [s1, s2].writeln;
}
