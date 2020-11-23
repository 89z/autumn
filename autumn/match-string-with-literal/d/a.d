import
   std.algorithm.searching,
   std.stdio;

void main() {
   auto s = "March";
   auto b = s.canFind("ar");
   b.writeln;
}
