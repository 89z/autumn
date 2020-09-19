import std.stdio, std.string;

void main() {
   auto s = "May June";
   // example 1
   auto a1 = s.split(" ");
   // example 2
   auto a2 = s.split;
   // print
   writeln(a1, a2);
}
