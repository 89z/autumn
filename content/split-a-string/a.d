import std.stdio, std.string;

void main() {
   auto s = "May June";
   // example A
   auto aA = s.split(" ");
   // example B
   auto aB = s.split;
   // print
   writeln(aA, aB);
}
