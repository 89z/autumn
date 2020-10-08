import std.stdio;

void main() {
   string[] a;
   // example 1
   a ~= "May";
   // example 2
   a ~= ["Jun", "Jul"];
   // print
   a.writeln;
}
