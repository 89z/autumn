import std.conv, std.stdio;

void main() {
   { // example 1
      auto s = "100";
      auto n = s.to!int;
      writeln(n == 100);
   }
   { // example 2
      auto s = "99.9";
      auto n = s.to!float;
      writeln(n == 99.9f);
   }
}
