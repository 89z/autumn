import std.stdio;

void main() {

   { // example 1
      auto s = r"March
April";
      s.writeln;
   }

   { // example 2
      auto s = `March
April`;
      s.writeln;
   }

}
