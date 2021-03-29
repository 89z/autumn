import std.array, std.stdio;

void main() {
   auto a = ["West", "East"];
   { // example 1
      auto s = a.join(",");
      writeln(s == "West,East");
   }
   { // example 2
      auto s = a.join;
      writeln(s == "WestEast");
   }
}
