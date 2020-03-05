import std.stdio;
void main() {
   const a1 = ["sun", "mon", "tue"];
   // string from front
   const s1 = a1[0];
   // string from back
   const s2 = a1[$ - 1];
   // array from front
   const a2 = a1[0 .. 2];
   // array from back
   const a3 = a1[$ - 2 .. $];
   // print
   writeln(s1, "\n", s2, "\n", a2, "\n", a3);
}
