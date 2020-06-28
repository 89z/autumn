import std.stdio;
void main() {
   auto s = "Sunday";
   // example 1
   auto s1 = s[0];
   s1.writeln;
   // example 2
   auto s2 = s[0 .. 2];
   s2.writeln;
}
