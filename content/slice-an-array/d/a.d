import std.stdio;

void main() {
   auto a = ["J", "u", "n", "e"];
   // example 1
   auto s1 = a[2];
   // example 2
   auto a2 = a[2 .. $];
   // print
   writeln(s1, a2);
}
