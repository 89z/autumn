import std.stdio;
void main() {
   auto s1 = "Sunday";
   auto s2 = s1[$ - 2 .. $];
   s2.writeln;
}
