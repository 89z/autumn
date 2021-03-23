import std.stdio;

void main() {
   string s;
   // quote
   auto s = r"one\two";
   auto s = `zero"one\two`;
   // slash
   auto s = r"one\two";
   auto s = `zero"one\two`;
   // newline
   auto s = r"March
   April";
   auto s = `March
   April`;
}
