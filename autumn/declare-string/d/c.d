import std.stdio;

void main() {
   // example 1
   auto s1 = r"one\two";
   // example 2
   auto s2 = `zero"one\two`;
   // print
   writeln(s1, ',', s2);
}
