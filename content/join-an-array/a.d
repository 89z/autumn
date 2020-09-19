import std.array, std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto s1 = a.join(",");
   // example 2
   auto s2 = a.join;
   // print
   writeln(s1 == "May,June" && s2 == "MayJune");
}
