import std.stdio;
void main() {
   int n1 = 0;
   string s1;
   if (n1 > 0) {
      s1 = "positive";
   } else if (n1 < 0) {
      s1 = "negative";
   } else {
      s1 = "zero";
   }
   s1.writeln;
}
