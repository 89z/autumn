import std.stdio;

void main() {
   string s;

   // slash
   s = "south\\north";
   s.writeln;

   // quote
   s = "south\"north";
   s.writeln;

   // newline
   s = "south
north";
   s.writeln;

}
