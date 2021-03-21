import std.stdio;

void main() {
   { // example 1
      int[string] m;
      m.writeln;
   }
   { // example 2
      auto m = ["month": 12, "day": 31];
      m.writeln;
   }
}
