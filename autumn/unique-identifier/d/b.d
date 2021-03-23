import std.conv, std.stdio;

void main() {
   string s;
   uint n;

   // example 1
   n = 0xFFFF_FFFF;
   s = n.to!string(36);
   writeln(s == "1Z141Z3");

   // example 2
   s = "1Z141Z3";
   n = s.to!uint(36);
   writeln(n == 0xFFFF_FFFF);

}
