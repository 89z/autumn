import std.stdio;

void main() {
   string s;

   // quote
   s = "south\"north";
   s.writeln;

   // slash
   s = "south\\north";
   s.writeln;

   // newline
   s = "south
north";
   s.writeln;

}
