import std.stdio;
void main() {
   int n = 10;
   string s;
   if (n > 12) {
      s = "Tue";
   } else if (n > 11) {
      s = "Mon";
   } else {
      s = "Sun";
   }
   s.writeln;
}
