import std.stdio;

void main() {
   auto a = ["M", "a", "r", "c", "h"];
   // example 1
   auto a1 = a[2 .. 4];
   // example 2
   auto a2 = a[2 .. $];
   // print
   writeln(a1, a2);
}
