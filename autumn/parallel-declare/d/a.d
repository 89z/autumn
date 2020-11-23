import
   std.stdio,
   std.typecons;

void main() {
   // example 1
   Tuple!(int, "a", int, "b") c = [9, 8];
   // example 2
   Tuple!(string, "a", int, "b") d = tuple("year", 2019);
   // print
   writeln(c.a == 9 && c.b == 8 && d.a == "year" && d.b == 2019);
}
