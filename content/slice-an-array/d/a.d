import std.stdio;

void main() {
   auto a = ["M", "a", "r", "c", "h"];
   // example 1
   auto s1 = a[2];
   // example 2
   auto a2 = a[2 .. 4];
   // example 3
   auto a3 = a[2 .. $];
   // print
   writeln(s1, a2, a3);
}
