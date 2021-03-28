import std.stdio;

void main() {
   auto s = "West";
   s ~= "East";
   writeln(s == "WestEast");
}
