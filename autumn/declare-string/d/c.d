import std.stdio;

void main() {
   { // example 1
      auto s = r"one\two";
      s.writeln;
   }
   { // example 2
      auto s = `zero"one\two`;
      s.writeln;
   }
}
