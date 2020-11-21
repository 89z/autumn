import std.stdio, std.typecons;

void main() {
   // example 1
   Tuple!(int, "a", int, "b") t = [8, 9];
   writeln(t.a == 8 && t.b == 9);
   // example 2
   Tuple!(string, "a", int, "b") u = tuple("year", 2019);
   writeln(u.a == "year" && u.b == 2019);
   // example 3
   string a = "year", b = "month";
   auto a = 9, b = 8;
   auto a = "year", b = 2019;
   // print
   writeln(a, ',', b);
}
