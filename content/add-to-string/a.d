import std.stdio;
void main() {
   // example 1
   auto s1 = "Sun";
   s1 = s1 ~ "day";
   // example 2
   auto s2 = "Sun";
   s2 ~= "day";
   // print
   [s1, s2].writeln;
}
