import std.algorithm, std.stdio;

void main() {
   { // example 1
      auto a = ["May", "June"];
      a.sort;
      a.writeln;
   }
   { // example 2
      auto a = [10, 9];
      a.sort;
      a.writeln;
   }
}
