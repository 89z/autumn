import std.stdio;

void main() {
   { // example 1
      int[] a = [10, 11];
      a.writeln;
   }
   { // example 2
      auto a = [10, 11];
      a.writeln;
   }
   { // example 3
      auto a = [[10, 11], [12, 13]];
      a.writeln;
   }
}
