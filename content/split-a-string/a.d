import std.stdio, std.string;
void main() {
   // example 1
   auto a1 = "Sun Mon".split(" ");
   // example 2
   auto a2 = "Sun Mon".split;
   // example 3
   auto a3 = `Sunday
Monday`.splitLines;
   // print
   writeln(a1, "\n", a2, "\n", a3);
}
