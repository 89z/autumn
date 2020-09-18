import std.stdio;

void main() {
   // example A
   auto sA = "May
June";
   // example B
   auto sB = `May
June`;
   // example C
   auto sC = q{May
June};
   // print
   [sA, sB, sC].writeln;
}
