import std.array, std.stdio;

void main() {
   auto s = "M a r c h";
   // example 1
   auto a1 = s.split(" ");
   // example 2
   auto a2 = s.split;
   // print
   writeln(a1, a2);
}
