import std.stdio;

void main() {
   auto s = "May";
   s ~= "June";
   writeln(s == "MayJune");
}
