import std.stdio;

void main() {
   // example 1
   auto s1 = r"sigma\tau";
   // example 2
   auto s2 = `sigma\tau`;
   // example 3
   auto s3 = `["sigma", "tau"]`;
   // print
   writeln(s1, ',', s2, ',', s3);
}
