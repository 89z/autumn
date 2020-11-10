import std.stdio;

void main() {
   string[] a;
   // example 1
   a ~= "March";
   // example 2
   a ~= ["April", "May"];
   // print
   a.writeln;
}
