import std.stdio;
void main() {
   auto a1 = ["Sun", "Mon", "Tue"];
   // string from front
   auto s1 = a1[0];
   // string from back
   auto s2 = a1[$ - 1];
   // array from front
   auto a2 = a1[0 .. 2];
   // array from back
   auto a3 = a1[$ - 2 .. $];
   // print
   writeln(s1, "\n", s2, "\n", a2, "\n", a3);
}
