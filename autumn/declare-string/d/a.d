import std.stdio;

void main() {
   auto s = q"[zero"one\two]";
   s.writeln;
}
