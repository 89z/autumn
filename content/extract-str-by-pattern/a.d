import std.algorithm, std.regex, std.stdio;
void main() {
   auto s1 = "Wednesday";
   auto s2 = "e.";
   auto a1 = s1.matchAll(s2.regex);
   a1.writeln;
}
