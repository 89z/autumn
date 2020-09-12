import std.stdio;
void main() {
   int n = 10;
   string s;
   if (n > 0) {
      s = "+";
   } else if (n < 0) {
      s = "-";
   } else {
      s = "zero";
   }
   writeln(s == "+");
}
