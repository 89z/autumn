import std.stdio, std.typecons;

void main() {
   { // example 1
      Tuple!(int, "m", int, "d") t = [12, 31];
      writeln(t.m == 12 && t.d == 31);
   }
   { // example 2
      Tuple!(string, "s", int, "n") t = tuple("month", 12);
      writeln(t.s == "month" && t.n == 12);
   }
}
