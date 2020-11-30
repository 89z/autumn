import
   std.stdio,
   std.typecons;

void main() {
   // example 1
   Tuple!(int, "m", int, "d") o1 = [12, 31];
   // example 2
   Tuple!(string, "s", int, "n") o2 = tuple("month", 12);
   // print
   writeln(o1.m == 12 && o2.n == 12);
}
