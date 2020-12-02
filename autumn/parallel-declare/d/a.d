import
   io = std.stdio,
   type = std.typecons: Tuple;

void main() {
   // example 1
   Tuple!(int, "m", int, "d") o1 = [12, 31];
   // example 2
   Tuple!(string, "s", int, "n") o2 = type.tuple("month", 12);
   // print
   io.writeln(o1.m == 12 && o2.n == 12);
}
