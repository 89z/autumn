import std.stdio, std.string;

void main() {
   auto s = "May June";
   // example 1
   auto a = s.split(" ");
   // example 2
   auto a2 = s.split;
   // print
   writeln(a, a2);
}
