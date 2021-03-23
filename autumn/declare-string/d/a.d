import std.stdio;

void main() {
   string s;

   // quote
   s = q"[south"north]";
   s.writeln;

   // slash
   s = q"[south\north]";
   s.writeln;

   // newline
   s = q"[south
north]";
   s.writeln;

}
