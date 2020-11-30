import std.stdio;

void main() {
   // example 1
   auto s1 = q"[from\to]";
   // example 2
   auto s2 = q"[["from", "to"]]";
   // print
   writeln(s1, s2);
}
