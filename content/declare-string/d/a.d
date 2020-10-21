import std.stdio;

void main() {
   // example 1
   string s1;
   // example 2
   string s2 = "one\\two";
   // exmaple 3
   auto s3 = "one\\two";
   // example 4
   auto s4 = `one\two`;
   // example 5
   auto s5 = r"one\two";
   // example 6
   auto s6 = q"[one\two]";
   // print
   [s1, s2, s3, s4, s5, s6].writeln;
}
