import std.array, std.stdio;

void main() {
   auto a = ["May", "June"];
   { // example 1
      auto s = a.join(",");
      writeln(s == "May,June");
   }
   { // example 2
      auto s = a.join;
      writeln(s == "MayJune");
   }
}
